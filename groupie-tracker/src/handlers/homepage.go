package handlers

import (
	"html/template"
	"net/http"
)

var homepage = template.Must(template.ParseFiles("homepage.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {

	err := homepage.ExecuteTemplate(w, "homepage.html", "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
