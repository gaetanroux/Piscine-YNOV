package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

var createpost = template.Must(template.ParseFiles("./html/createpost.html"))

// AddPost adds posts to the database
func AddPost(title string, body string, userID string) {

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	request := `
		INSERT INTO posts(title, contenu, user_id, created_at)
		VALUES (?, ?, ?, datetime('now','localtime'))
	`
	_, err = db.Exec(request, title, body, userID)
}

// CreatePost handle createpost tab
func CreatePost(w http.ResponseWriter, r *http.Request) {

	files := []string{"./html/createpost.html", "./html/base.html"}

	user := GetAuthentificatedUser(r)
	title := r.FormValue("title")
	body := r.FormValue("body")
	userID := user.ID

	if user.Connected == true && title != "" && body != "" {

		AddPost(title, body, userID)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
