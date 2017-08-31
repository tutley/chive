package handlers

import (
	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	//"log"
	"net/http"
	//	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
)

// Index serves as the anchor for all the handlers based on top-level routes
type Index struct {
	RespFormat  string
	TemplateBox *rice.Box
	Db          *mgo.Database
}

// Routes creates a REST router for the index resource
func (rs Index) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	//r.Use(middleware.WithValue("respFormat", rs.RespFormat))

	r.Get("/", rs.Home) // GET /

	return r
}

// Home grabs the home page
func (rs Index) Home(w http.ResponseWriter, r *http.Request) {
	//	c := struct{}{Title: "chive", Author: "Tom Utley"}

	w.Write([]byte("HOME PAGE"))
	// This has to change to use the templates from rice
	//	err := templates.ExecuteTemplate(w, "index", c)
	//	if err != nil {
	//		log.Println(err)
	//	}
}
