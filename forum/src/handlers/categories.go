package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var categories = template.Must(template.ParseFiles("./html/categories.html"))

// Categories handle categories tab
func Categories(w http.ResponseWriter, r *http.Request) {

	files := []string{"./html/categories.html", "./html/base.html"}

	user := GetAuthentificatedUser(r)

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
