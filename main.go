package main

import (
	"context"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type key int

var dbKey key = 100000

// getDb grabs the mgo database from the context
func getDb(ctx context.Context) *mgo.Database {
	return ctx.Value(dbKey).(*mgo.Database)
}

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
	mongoMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// setup the mgo connection
			session, err := mgo.Dial("mongodb://" + dbURL)

			if err != nil {
				log.Println("DB Connect error: ", err)
				http.Error(w, "Unable to connect to database", 500)
			}

			reqSession := session.Clone()
			defer reqSession.Close()
			db := reqSession.DB(dbName)
			ctx := context.WithValue(r.Context(), helpers.DbKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	r := chi.NewRouter()

	// Use Chi built-in middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mount database
	r.Use(mongoMiddleware)

	// Setup routes
	r.Get("/", getIndexHandler)

	// Rather than doing server-side rendering, we are using templates to populate
	// all of the meta tags in the header of the web pages, so that if a scraper or
	// a search engine tries to access a link, they will get all of the page information
	// even if they can't execute javascript. So in this routing section, we duplicate
	// all of the routes but use the same underlying router to handle them.
	r.Mount("/examples", exampleRoutes("template"))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/examples", exampleRoutes("json"))
	})

	// This serves the static files
	//fileServer(r, "/dist", assetFS())

	serveAddr := ":" + strconv.Itoa(serverPort)
	log.Println("dhcpportal Server listening on: ", strconv.Itoa(serverPort))
	http.ListenAndServe(serveAddr, r)
	// TODO: convert to HTTPS
	// For publishing, get a cert with LetsEncrypt
	// https://golang.org/pkg/net/http/#ListenAndServeTLS
}

// exampleRoutes is the router for example objects
func exampleRoutes(string respFormat) chi.Router {
	r := chi.NewRouter()
	rf := "json"
	rf = respFormat
	r.Use(middleware.WithValue("respFormat", rf))
	// TODO: add JWT authentication checking

	r.Get("/", getExampleList)  // GET /examples
	r.Post("/", postNewExample) // POST /examples

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getExample)       // GET /examples/{id}
		r.Put("/", putExample)       // PUT /examples/{id}
		r.Delete("/", deleteExample) // DELETE /examples/{id}
	})

	return r
}

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
