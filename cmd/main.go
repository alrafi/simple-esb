package main

import (
	"log"
	"net/http"
	"simple-esb/internal/handler"
)

func main() {
	log.Print("Simple ESB Server run on port 8080...")

	http.HandleFunc("/continents", handler.GetContinentList)
	http.HandleFunc("/capitalcity", handler.GetCapitalCity)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
