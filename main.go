package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"gopkg.in/mgo.v2"

	// BE sure to change these for each new project
	"github.com/tutley/chive/handlers"
)

func main() {

	serverPort := 3333
	// server port is a command line object so that you can reverse proxy with something
	// like NGINX and have a bunch of web apps running on one server
	var sPort string
	flag.StringVar(&sPort, "p", "3333", "The TCP Port for serving this webpage - default 3333")
	flag.Parse()

	sp, err := strconv.Atoi(sPort)
	if err != nil {
		log.Println("Error: port input at the command line failed to be applied: ", sPort)
	} else {
		serverPort = sp
	}

	// All of the other variables will be compiled into the app so set them here
	dbURL := "localhost"
	dbName := "chive"

	// Init the Database
	session, err := mgo.Dial("mongodb://" + dbURL)
	if err != nil {
		log.Fatal("DB Connect error: ", err)
	}

	defer session.Close()
	db := session.DB(dbName)

	// init the templates
	templateBox, err := rice.FindBox("chive-templates")
	if err != nil {
		log.Fatal(err)
	}
	handlers.SetTemplateBox(templateBox)

	r := chi.NewRouter()
	// Use Chi built-in middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultCompress)
	// Setup routes
	r.Mount("/", handlers.Index{
		RespFormat: "template",
		Db:         db}.Routes())

	// Rather than doing server-side rendering, we are using templates to populate
	// all of the meta tags in the header of the web pages, so that if a scraper or
	// a search engine tries to access a link, they will get all of the page information
	// even if they can't execute javascript. So in this routing section, we duplicate
	// all of the routes but use the same underlying router to handle them.

	r.Mount("/examples", handlers.Examples{
		RespFormat: "template",
		Db:         db}.Routes())

	r.Route("/api", func(r chi.Router) {

		// TODO: Narrow this down so not everyone in the world can make API requests
		cors := cors.New(cors.Options{
			// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(cors.Handler)
		r.Mount("/examples", handlers.Examples{
			RespFormat: "json",
			Db:         db}.Routes())
	})

	// This serves the static files
	// There is probably a better way to do this, but for now I'm manually serving the service worker files
	box := rice.MustFindBox("chive-dist")
	sw, _ := box.String("service-worker.js")
	r.Get("/service-worker.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(sw))
	})
	wbjs, _ := box.String("workbox-sw.prod.v2.1.2.js")
	r.Get("/workbox-sw.prod.v2.1.2.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(wbjs))
	})
	wbsm, _ := box.String("workbox-sw.prod.v2.1.2.js.map")
	r.Get("/workbox-sw.prod.v2.1.2.js.map", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(wbsm))
	})
	// OK now mount the static folder from within foundirl-dist
	distFileServer := http.StripPrefix("/", http.FileServer(box.HTTPBox()))
	r.Mount("/static/", distFileServer)

	serveAddr := ":" + strconv.Itoa(serverPort)
	log.Println("chive Server listening on: ", strconv.Itoa(serverPort))
	http.ListenAndServe(serveAddr, r)
}
