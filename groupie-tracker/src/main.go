package main

import (
	"net/http"

	"./handlers"
)

func main() {
	fss := http.FileServer(http.Dir("./musique"))
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/musique/", http.StripPrefix("/musique/", fss))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/artists", handlers.ArtistsFunc)
	http.HandleFunc("/artist", handlers.ArtistFunc)
	http.HandleFunc("/concert", handlers.Concert)
	http.ListenAndServe(":8080", nil)
}
