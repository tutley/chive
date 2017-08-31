package main

import (
	//"context"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strconv"

	// BE sure to change these for each new project
	"github.com/tutley/chive/handlers"
	//"github.com/tutley/chive/models"
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

	r := chi.NewRouter()

	// Use Chi built-in middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Setup routes

	r.Mount("/", handlers.Index{
		RespFormat:  "template",
		TemplateBox: templateBox,
		Db:          db}.Routes())

	// Rather than doing server-side rendering, we are using templates to populate
	// all of the meta tags in the header of the web pages, so that if a scraper or
	// a search engine tries to access a link, they will get all of the page information
	// even if they can't execute javascript. So in this routing section, we duplicate
	// all of the routes but use the same underlying router to handle them.

	r.Mount("/examples", handlers.Examples{
		RespFormat:  "template",
		TemplateBox: templateBox,
		Db:          db}.Routes())

	r.Route("/api", func(r chi.Router) {
		r.Mount("/examples", handlers.Examples{
			RespFormat:  "json",
			TemplateBox: templateBox,
			Db:          db}.Routes())
	})

	// This serves the static files
	r.Mount("/dist", http.FileServer(rice.MustFindBox("chive-dist").HTTPBox()))

	serveAddr := ":" + strconv.Itoa(serverPort)
	log.Println("dhcpportal Server listening on: ", strconv.Itoa(serverPort))
	http.ListenAndServe(serveAddr, r)
	// TODO: convert to HTTPS
	// For publishing, get a cert with LetsEncrypt
	// https://golang.org/pkg/net/http/#ListenAndServeTLS
}
