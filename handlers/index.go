package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	//	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
)

// Index serves as the anchor for all the handlers based on top-level routes
type Index struct {
	RespFormat string
	Db         *mgo.Database
}

// Routes creates a REST router for the index resource
func (rs Index) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..
	//r.Use(middleware.WithValue("respFormat", rs.RespFormat))

	r.Get("/", rs.Home) // GET /
	r.NotFound(rs.Home)
	return r
}

// Home grabs the home page
func (rs Index) Home(w http.ResponseWriter, r *http.Request) {
	mets := genMetas()
	sendTemplate(mets, w, r)
}
