package main

import (
	"encoding/json"
	"html/template"
	//"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strings"
	"time"
)

// use go.rice to import the templates
var templates *template.Template

// PrettyTime is used in the html template to make the time pretty
// func PrettyTime(t time.Time) string {
// 	pt := t.Format("01/02/2006 3:04 PM MST")
// 	return pt
// }

// GetIndexHandler grabs the home page
func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	c := struct{}{Title: "chive", Author: "Tom Utley"}

	// This has to change to use the templates from rice
	err := templates.ExecuteTemplate(w, "index", c)
	if err != nil {
		log.Println(err)
	}
}
