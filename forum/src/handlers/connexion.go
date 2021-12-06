package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	Connected bool
	Cookie    string
	ID        string
	Username  string
	FirstName string
	LastName  string
	Email     string
}

var connexion = template.Must(template.ParseFiles("./html/connexion.html"))

// GetAuthentificatedUser get the authentificated user
func GetAuthentificatedUser(r *http.Request) User {

	user := User{}

	c, err := r.Cookie("cookie_session")
	if err != nil {
		user.Connected = false
		user.Cookie = ""
	} else {
		user.Connected = true
		user.Cookie = c.Name
	}

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	if user.Connected == true {

		idreq := `
			SELECT user_id FROM user_session WHERE cookie_expiration = ?;
		`

		id := ""

		_ = db.QueryRow(idreq, c.Value).Scan(&id)
		if err != nil {
			fmt.Print("Erreur")
		}

		user.ID = id

		getUsername := `
		SELECT username FROM users WHERE id IS ?;
		`

		getFirstName := `
		SELECT first_name FROM users WHERE id IS ?;
		`

		getLastName := `
		SELECT last_name FROM users WHERE id IS ?;
		`

		getEmail := `
		SELECT email FROM users WHERE id IS ?;
		`

		_ = db.QueryRow(getUsername, user.ID).Scan(&user.Username)
		if err != nil {
			fmt.Print("Erreur")
		}

		_ = db.QueryRow(getFirstName, user.ID).Scan(&user.FirstName)
		if err != nil {
			fmt.Print("Erreur")
		}

		_ = db.QueryRow(getLastName, user.ID).Scan(&user.LastName)
		if err != nil {
			fmt.Print("Erreur")
		}

		_ = db.QueryRow(getEmail, user.ID).Scan(&user.Email)
		if err != nil {
			fmt.Print("Erreur")
		}

	} else {

		user.ID = ""
	}

	return user
}

// CheckPasswordHash checks if the password match with the hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func connected(email string, password string) bool {

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	mail := `
		SELECT email FROM users WHERE email = ?;
	`

	checkmail := `
		SELECT password FROM users WHERE email = ?;
	`

	mailexist := ""
	passw := ""

	_ = db.QueryRow(mail, email).Scan(&mailexist)
	if err != nil {
		fmt.Print("Erreur")
		return false
	}

	_ = db.QueryRow(checkmail, email).Scan(&passw)
	if err != nil {
		fmt.Print("Erreur")
		return false
	}

	hash := CheckPasswordHash(password, passw)

	if email != "" {

		if email == mailexist {

			if password != "" {

				if hash == true {

					return true
				}

				fmt.Print("Mot de passe incorrect")
				return false
			}
		}

		fmt.Print("Adresse mail incorrect !")
		return false
	}

	return false
}

// CreateCookie creates a new cookie
func CreateCookie(w http.ResponseWriter, r *http.Request, email string) {

	c, err := r.Cookie("cookie_session")
	if err != nil {
		sID := uuid.NewV4()
		if err != nil {
			fmt.Print(err)
		}
		c = &http.Cookie{
			Name:    "cookie_session",
			Value:   sID.String(),
			Expires: time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, c)
	}

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	idreq := `
		SELECT id FROM users WHERE email = ?;
	`

	id := ""

	_ = db.QueryRow(idreq, email).Scan(&id)
	if err != nil {
		fmt.Print("Erreur")
	}

	request := `
	INSERT INTO user_session(cookie_expiration, user_id)
	VALUES (?, ?)
	`

	_, err = db.Exec(request, c.Value, id)
	if err != nil {
		fmt.Print(err)
	}
}

// Connexion handle connexion tab
func Connexion(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("pass")

	db, err := sql.Open("sqlite3", "./Database/database.db")
	if err != nil {
		panic(err)
	}

	idreq := `
		SELECT id FROM users WHERE email = ?;
	`

	id := ""

	_ = db.QueryRow(idreq, email).Scan(&id)
	if err != nil {
		fmt.Print("Erreur")
	}

	var user User

	if connected(email, password) == true && user.Connected == false {

		CreateCookie(w, r, email)
		user.Connected = true
		user.Cookie = "cookie_session"
		user.ID = id
		fmt.Print("Cookie créé !\n")
		fmt.Print("Vous êtes connectés à ce super site !\n")
		http.Redirect(w, r, "/", 302)
	}

	files := []string{"./html/connexion.html", "./html/base.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Print(err)
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
