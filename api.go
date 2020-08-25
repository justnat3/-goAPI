package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

//Generic is a model for rec
type Generic struct {
	Response string `json:"Response"`
}

//Rec implements Generics
var Rec []Generic
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//Generate random string
func random(num int) string {
	b := make([]rune, num)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//EndPoint for blank string
func getNameWithoutPrefix(writer http.ResponseWriter, req *http.Request) {
	Rec := append(Rec, Generic{Response: random(10)})
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Rec)
}

func main() {

	router := mux.NewRouter()
	port := ":5000"
	println("Listening on port", port)

	randstring := random(10)
	println(randstring)

	router.HandleFunc("/nameGen", getNameWithoutPrefix).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
