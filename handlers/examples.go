package handlers

import (
	"encoding/json"
	"errors"
	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tutley/chive/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
)

// Examples serves as the anchor for all the handlers based on examples routes
type Examples struct {
	RespFormat  string
	TemplateBox *rice.Box
	Db          *mgo.Database
}

var masterTpl *template.Template

// Routes creates a REST router for the examples resource
// /examples
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

	// prepare templates
	indexString, err := rs.TemplateBox.String("index.tpl")
	headerString, err := rs.TemplateBox.String("header.tpl")
	footerString, err := rs.TemplateBox.String("footer.tpl")

	// parse and execute the template
	tmpl, err := template.New("index").Parse(headerString)
	tmpl2, err := template.Must(tmpl.Clone()).Parse(footerString)
	tmpl3, err := template.Must(tmpl2.Clone()).Parse(indexString)
	masterTpl = tmpl3

	if err != nil {
		log.Fatal(err)
	}

	return r
}

// List lists all the examples
func (rs Examples) List(w http.ResponseWriter, r *http.Request) {
	respFormat := r.Context().Value("respFormat").(string)

	// Grab examples from DB
	examples, err := models.ListExamples(rs.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check the response format
	if respFormat == "json" {
		js, err := json.Marshal(examples)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		c := struct {
			Title string
		}{
			Title: "Chive - List of Examples",
		}
		err := masterTpl.ExecuteTemplate(w, "index", c)
		if err != nil {
			log.Println(err)
		}
	}
}

// Create will make a new example
func (rs Examples) Create(w http.ResponseWriter, r *http.Request) {
	respFormat := r.Context().Value("respFormat").(string)

	// check the response format
	if respFormat != "json" {
		// We don't allow creates from outside of json
		err := errors.New("You can't do that.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// go ahead and push it to the db and return success or error
	example := models.Example{}

	err := json.NewDecoder(r.Body).Decode(&example)
	if err != nil {
		log.Println("error parsing json request ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	example.ID = bson.NewObjectId()

	err = example.Insert(rs.Db)
	if err != nil {
		log.Println("error inserting example: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(example)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Get will get an example
func (rs Examples) Get(w http.ResponseWriter, r *http.Request) {
	respFormat := r.Context().Value("respFormat").(string)

	// Grab example from DB
	id := chi.URLParam(r, "id")
	example, err := models.FindExample(id, rs.Db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Println("Error finding example: ", err)
		return
	}

	// check the response format
	if respFormat == "json" {
		js, err := json.Marshal(example)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error encoding example into json: ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		var t string
		t = example.Title + " - Chive"
		c := struct {
			Title string
		}{
			Title: t,
		}
		err := masterTpl.ExecuteTemplate(w, "index", c)
		if err != nil {
			log.Println(err)
		}
	}

}

// Update will update an example
func (rs Examples) Update(w http.ResponseWriter, r *http.Request) {
	respFormat := r.Context().Value("respFormat").(string)

	// check the response format
	if respFormat != "json" {
		// We don't allow creates from outside of json
		err := errors.New("You can't do that.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := chi.URLParam(r, "id")

	// go ahead and push it to the db and return success or error
	example := models.Example{}

	err := json.NewDecoder(r.Body).Decode(&example)
	if err != nil {
		log.Println("error parsing json request ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	example.ID = bson.ObjectIdHex(id)

	err = example.Update(rs.Db)
	if err != nil {
		log.Println("error update example: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(204)
	w.Write([]byte("\n"))
}

// Delete will delete an example
func (rs Examples) Delete(w http.ResponseWriter, r *http.Request) {
	respFormat := r.Context().Value("respFormat").(string)

	// check the response format
	if respFormat != "json" {
		// We don't allow creates from outside of json
		err := errors.New("You can't do that.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := chi.URLParam(r, "id")

	example := models.Example{ID: bson.ObjectIdHex(id)}

	e := example.Delete(rs.Db)
	if e != nil {
		log.Println("error deleting example: ", e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
	w.Write([]byte("\n"))
}
