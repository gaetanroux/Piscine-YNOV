package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var homepage = template.Must(template.ParseFiles("./html/homepage.html"))

// Homepage struct
type Homepage struct {
	Connected bool
	Cookie    string
	ID        string
	FourPosts [4]OnePost
}

// OnePost struct
type OnePost struct {
	UserRecent  string
	DateRecent  string
	TitleRecent string
	BodyRecent  string
}

// HomePage handle homepage tab
func HomePage(w http.ResponseWriter, r *http.Request) {

	files := []string{"./html/homepage.html", "./html/base.html"}

	user := GetAuthentificatedUser(r)

	var accueil Homepage
	var TabPosts [4]OnePost

	accueil.Connected = user.Connected
	accueil.Cookie = user.Cookie
	accueil.ID = user.ID

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	var allDates string
	fourDates := `
	SELECT group_concat(created_at) FROM (SELECT created_at FROM posts ORDER BY created_at DESC LIMIT 4)
	`
	_ = db.QueryRow(fourDates).Scan(&allDates)
	if err != nil {
		fmt.Print("Erreur")
	}

	s := strings.Split(allDates, ",")

	for i := 0; i < len(s); i++ {

		TabPosts[i].DateRecent = s[i]

		getRecentTitle := `
		SELECT title FROM posts WHERE created_at IS ?;
		`

		_ = db.QueryRow(getRecentTitle, &TabPosts[i].DateRecent).Scan(&TabPosts[i].TitleRecent)
		if err != nil {
			fmt.Print("Erreur")
		}

		getRecentBody := `
		SELECT contenu FROM posts WHERE created_at IS ?;
		`

		_ = db.QueryRow(getRecentBody, &TabPosts[i].DateRecent).Scan(&TabPosts[i].BodyRecent)
		if err != nil {
			fmt.Print("Erreur")
		}

		getRecentID := `
		SELECT user_id FROM posts WHERE created_at IS ?;
		`

		recentID := ""

		_ = db.QueryRow(getRecentID, &TabPosts[i].DateRecent).Scan(&recentID)
		if err != nil {
			fmt.Print("Erreur")
		}

		getRecentUser := `
		SELECT username FROM users WHERE id IS ?;
		`

		_ = db.QueryRow(getRecentUser, &recentID).Scan(&TabPosts[i].UserRecent)
		if err != nil {
			fmt.Print("Erreur")
		}
	}

	// getRecentDate := `
	// SELECT MAX (created_at) AS "Max Date"
	// FROM posts;
	// `

	// _ = db.QueryRow(getRecentDate).Scan(&TabPosts[0].DateRecent)
	// if err != nil {
	// 	fmt.Print("Erreur")
	// }

	// getRecentTitle := `
	// SELECT title FROM posts WHERE created_at IS ?;
	// `

	// _ = db.QueryRow(getRecentTitle, &TabPosts[0].DateRecent).Scan(&TabPosts[0].TitleRecent)
	// if err != nil {
	// 	fmt.Print("Erreur")
	// }

	// getRecentBody := `
	// SELECT contenu FROM posts WHERE created_at IS ?;
	// `

	// _ = db.QueryRow(getRecentBody, &TabPosts[0].DateRecent).Scan(&TabPosts[0].BodyRecent)
	// if err != nil {
	// 	fmt.Print("Erreur")
	// }

	// getRecentID := `
	// SELECT user_id FROM posts WHERE created_at IS ?;
	// `

	// recentID := ""

	// _ = db.QueryRow(getRecentID, &TabPosts[0].DateRecent).Scan(&recentID)
	// if err != nil {
	// 	fmt.Print("Erreur")
	// }

	// getRecentUser := `
	// SELECT username FROM users WHERE id IS ?;
	// `

	// _ = db.QueryRow(getRecentUser, &recentID).Scan(&TabPosts[0].UserRecent)
	// if err != nil {
	// 	fmt.Print("Erreur")
	// }

	accueil.FourPosts = TabPosts

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
	}

	err = tmpl.Execute(w, accueil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
