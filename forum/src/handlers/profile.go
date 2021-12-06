package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

// ProfileUser struct
type ProfileUser struct {
	Connected   bool
	Cookie      string
	ID          string
	UserRecent  string
	DateRecent  string
	TitleRecent string
	BodyRecent  string
	Username    string
	FirstName   string
	LastName    string
	Email       string
}

var profile = template.Must(template.ParseFiles("./html/profile.html"))

// Profile handle profile tab
func Profile(w http.ResponseWriter, r *http.Request) {

	files := []string{"./html/profile.html", "./html/base.html"}

	user := GetAuthentificatedUser(r)

	var profil ProfileUser

	profil.Connected = user.Connected
	profil.Cookie = user.Cookie
	profil.ID = user.ID

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	getRecentDate := `
	SELECT MAX (created_at) AS "Max Date" 
	FROM posts WHERE user_id IS ?;
	`

	_ = db.QueryRow(getRecentDate, profil.ID).Scan(&profil.DateRecent)
	if err != nil {
		fmt.Print("Erreur")
	}

	getRecentTitle := `
	SELECT title FROM posts WHERE created_at IS ?; 
	`

	_ = db.QueryRow(getRecentTitle, &profil.DateRecent).Scan(&profil.TitleRecent)
	if err != nil {
		fmt.Print("Erreur")
	}

	getRecentBody := `
	SELECT contenu FROM posts WHERE created_at IS ?; 
	`

	_ = db.QueryRow(getRecentBody, &profil.DateRecent).Scan(&profil.BodyRecent)
	if err != nil {
		fmt.Print("Erreur")
	}

	getRecentID := `
	SELECT user_id FROM posts WHERE created_at IS ?; 
	`

	recentID := ""

	_ = db.QueryRow(getRecentID, &profil.DateRecent).Scan(&recentID)
	if err != nil {
		fmt.Print("Erreur")
	}

	getRecentUser := `
	SELECT username FROM users WHERE id IS ?;
	`

	_ = db.QueryRow(getRecentUser, &recentID).Scan(&profil.UserRecent)
	if err != nil {
		fmt.Print("Erreur")
	}

	profil.Username = user.Username
	profil.FirstName = user.FirstName
	profil.LastName = user.LastName
	profil.Email = user.Email

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
	}

	err = tmpl.Execute(w, profil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
