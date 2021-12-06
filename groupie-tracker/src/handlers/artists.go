package handlers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var artistes = template.Must(template.ParseFiles("artistes.html"))

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func ArtistsFunc(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var artists []Artist

	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Println(err)
	}

	ts, err := template.ParseFiles("artistes.html")
	if err != nil {
		log.Println(err)
	}

	err = ts.Execute(w, artists)
	if err != nil {
		log.Println(err)
	}
}
