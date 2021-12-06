package main

import (
	"net/http"
	"html/template"
	"./ASCIIART"
)

// Valeur renvoyée à la page web
type Page struct {
	Valeur string
}

var templates = template.Must(template.ParseFiles("index.html"))

// Permet la gestion de la page web
func asciiFunc(w http.ResponseWriter, r *http.Request) {

	// Récupération des données du formulaire nécessaires au lancement de la fonction PrintAscii
	text := r.FormValue("text")
	arr := []byte(text)
    nbrLigne := 1		
    font := r.FormValue("font")		

	// Effectue la fonction PrintAscii avec les données récupérées et prépare à l'affichage
	print := ASCIIART.PrintAscii(arr, nbrLigne, font)
	p := &Page{Valeur: print}

	// Affiche/actualise de la page web 
	err := templates.ExecuteTemplate(w, "index.html", p)
 	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	// Permet la création du serveur
	http.HandleFunc("/ascii-art", asciiFunc)
	// Mise en écoute sur le port 9999
	http.ListenAndServe(":9999", nil)
}
