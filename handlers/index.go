package handlers

import (
	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
	"html/template"
	"log"
	"net/http"
	//	"github.com/go-chi/chi/middleware"
	"gopkg.in/mgo.v2"
	"os"
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
	c := struct {
		Title  string
		Author string
	}{
		Title:  "Chive",
		Author: "Tom Utley",
	}

	// get file contents as string
	indexString, err := rs.TemplateBox.String("index.tpl")
	headerString, err := rs.TemplateBox.String("header.tpl")
	footerString, err := rs.TemplateBox.String("footer.tpl")

	// parse and execute the template
	tmpl, err := template.New("index").Parse(headerString)
	tmpl2, err := template.Must(tmpl.Clone()).Parse(footerString)
	tmpl3, err := template.Must(tmpl2.Clone()).Parse(indexString)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl3.ExecuteTemplate(os.Stdout, "index", c)
	err = tmpl3.ExecuteTemplate(w, "index", c)
	if err != nil {
		log.Println(err)
	}

}
