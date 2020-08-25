package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Roll is model for sushi
type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//init rolls var as a slice
var rolls []Roll

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
