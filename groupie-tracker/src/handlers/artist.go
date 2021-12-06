package handlers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var artiste = template.Must(template.ParseFiles("artiste.html"))

func ArtistFunc(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + r.FormValue("id"))
	if err != nil {

		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		log.Fatalln(err)
	}

	var artist Artist

	err = json.Unmarshal(body, &artist)
	if err != nil {
		log.Println(err)
	}

	ts, err := template.ParseFiles("artiste.html")
	if err != nil {
		log.Println(err)
	}

	err = ts.Execute(w, &artist)
	if err != nil {
		log.Println(err)
	}
}
