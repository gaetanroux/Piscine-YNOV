package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Comment to justify the import
	"golang.org/x/crypto/bcrypt"
)

var inscription = template.Must(template.ParseFiles("./html/inscription.html"))

// HashPassword crypts a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// createUser creates a new user
func createUser(username string, firstname string, lastname string, email string, pass string, confirmpass string, w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	hashpass, _ := HashPassword(pass)

	if pass != confirmpass {

		print("Mots de passe différents")
	} else {

		request := `
			INSERT INTO users(username, first_name, last_name, email, password)
			VALUES (?, ?, ?, ?, ?)
		`

		_, err = db.Exec(request, username, firstname, lastname, email, hashpass)
		if email != "" && err != nil {
			fmt.Print("Adresse email déjà utilisée.")
		}
	}
}

// Inscription handle inscription tab
func Inscription(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	confirmpass := r.FormValue("confirm_password")

	createUser(username, firstname, lastname, email, pass, confirmpass, w, r)

	files := []string{"./html/inscription.html", "./html/base.html"}

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
