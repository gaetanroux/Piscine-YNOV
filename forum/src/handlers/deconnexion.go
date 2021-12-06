package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

var deconnexion = template.Must(template.ParseFiles("./html/deconnexion.html"))

// DeleteCookie deletes cookie
func DeleteCookie(user User) {

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	deletereq := `
	DELETE FROM user_session WHERE user_id = ?;
	`

	delete := ""

	_ = db.QueryRow(deletereq, user.ID).Scan(&delete)
	if err != nil {
		fmt.Print("Erreur")
	}
}

// Deconnexion handle deconnexion tab
func Deconnexion(w http.ResponseWriter, r *http.Request) {

	user := GetAuthentificatedUser(r)
	answer := r.FormValue("confirmation")

	if answer == "Oui" {

		c := &http.Cookie{
			Name:   "cookie_session",
			Value:  "Disconnected",
			MaxAge: -1,
		}
		http.SetCookie(w, c)
		DeleteCookie(user)
		fmt.Print("Cookie deleted !")
		http.Redirect(w, r, "/", 302)
	} else if answer == "Non" {

		fmt.Print("Vous êtes toujours connecté.")
		http.Redirect(w, r, "/", 302)
	}

	files := []string{"./html/deconnexion.html", "./html/base.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
