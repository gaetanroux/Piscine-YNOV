package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Locations struct {
	Index []struct {
		ID        int      `json:"id"` 
		Locations []string `json:"locations"`
		Dates string `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type City struct {
    Name    string
	Country string
	Slug string
}

var concert = template.Must(template.ParseFiles("concert.html"))

func Concert(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Print("1 ")
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("2 ")
		log.Fatalln(err)
	}

	var location Locations
	
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Print("3 ")
		log.Println(err)
	}

	temp := []string{}
	dates := []Dates{}
	date := Dates{}

	for _, tmp := range location.Index {

		for _, c := range tmp.Locations {

			temp = append(temp, c)
		}
		resp, err := http.Get(tmp.Dates)
		if err != nil {
			fmt.Print("4 ")
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print("5 ")
			log.Fatalln(err)
		}

		err = json.Unmarshal(body, &date)
		if err != nil {
			fmt.Print("6 ")
			log.Println(err)
		}

		dates = append(dates, date)
	}

	city := []City{}

	for _, location := range temp {

		tempo := strings.Split(location, "-")
		city = append(city, City{Name : tempo[0], Country : tempo[1]})
	}
	
	fmt.Print(location.Index[20])
	fmt.Print(dates[20])

	ts, err := template.ParseFiles("concert.html")
	if err != nil {
		fmt.Print("7 ")
		log.Println(err)
	}

	err = ts.Execute(w, location)
	if err != nil {
		fmt.Print("8 ")
		log.Println(err)
	}
}