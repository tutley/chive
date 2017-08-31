package handlers

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
)

// Examples serves as the anchor for all the handlers based on examples routes
type Examples struct {
	RespFormat  string
	TemplateBox *rice.Box
	Db          *mgo.Database
}

// Routes creates a REST router for the examples resource
func (rs Examples) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	r.Use(middleware.WithValue("respFormat", rs.RespFormat))
	// TODO: add JWT authentication checking

	r.Get("/", rs.List)    // GET /examples - read a list of examples
	r.Post("/", rs.Create) // POST /examples - create a new example and persist it

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /examples/{id} - read a single example by :id
		r.Put("/", rs.Update)    // PUT /examples/{id} - update a single example by :id
		r.Delete("/", rs.Delete) // DELETE /examples/{id} - delete a single example by :id
	})

	return r
}

// List lists all the examples
func (rs Examples) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("examples list of stuff.."))
}

// Create will make a new example
func (rs Examples) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("examples create"))
}

// Get will get an example
func (rs Examples) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("example get"))
}

// Update will update an example
func (rs Examples) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("example update"))
}

// Delete will delete an example
func (rs Examples) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("example delete"))
}
