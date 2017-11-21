package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

// this is where we setup a helper function to use the templates to respond to direct URL requests
// so if a person or crawler goes to "/" they get the default app with generic descriptions
// if they go to "/examples/1234" it will replace the meta tags with info specific to example 1234

var masterTpl *template.Template
var templateBox *rice.Box

type siteMetas struct {
	SiteName        string
	Title           string // Title of this asset
	Description     string
	ImageURL        string
	URL             string
	TwitterUsername string // including the @
	Type            string // website or article
}

func genMetas() siteMetas {
	m := siteMetas{
		SiteName:    "Chive",
		Title:       "Welcome to Chive",
		Description: "A full stack website boilerplate written in Go and Vue.",
		ImageURL:    "https://i.imgur.com/PComu8U.jpg",
		Type:        "website",
	}
	return m
}

// SetTemplateBox is a setter to allow the main program to init the templates for the handlers package
func SetTemplateBox(t *rice.Box) {
	templateBox = t
}

// sendTemplate is a handlerfunc that will take the template metas and produce an HTML response
func sendTemplate(metas siteMetas, w http.ResponseWriter, r *http.Request) {
	// get file contents as string
	indexString, err := templateBox.String("index.tpl")
	headerString, err := templateBox.String("header.tpl")
	footerString, err := templateBox.String("footer.tpl")

	// TODO: change yourdomain.com to your actual domain
	metas.URL = "https://yourdomain.com" + r.URL.Path

	// parse and execute the template
	tmpl, err := template.New("index").Parse(headerString)
	tmpl2, err := template.Must(tmpl.Clone()).Parse(footerString)
	tmpl3, err := template.Must(tmpl2.Clone()).Parse(indexString)
	if err != nil {
		log.Println("Error parsing templates ", err)
		http.Error(w, "We messed up and the entire site is broke.", http.StatusInternalServerError)
		return
	}
	err = tmpl3.ExecuteTemplate(w, "index", metas)
	if err != nil {
		http.Error(w, "Oh boy this is bad.", http.StatusInternalServerError)
		log.Println("Error executing template ", err)
		return
	}
}
