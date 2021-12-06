package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var post = template.Must(template.ParseFiles("./html/post.html"))

// Post handle post tab
func Post(w http.ResponseWriter, r *http.Request) {

	files := []string{"./html/post.html", "./html/base.html"}

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
