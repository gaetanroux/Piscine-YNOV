package main

import (
	"FORUM/src/handlers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./css"))
	fss := http.FileServer(http.Dir("./images"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", handlers.HomePage)
	http.Handle("/images/", http.StripPrefix("/images/", fss))
	http.HandleFunc("/categories", handlers.Categories)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/createpost", handlers.CreatePost)
	http.HandleFunc("/inscription", handlers.Inscription)
	http.HandleFunc("/post", handlers.Post)
	http.HandleFunc("/profile", handlers.Profile)
	http.HandleFunc("/deconnexion", handlers.Deconnexion)
	http.ListenAndServe(":8080", nil)

}
