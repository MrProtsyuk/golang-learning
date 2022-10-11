package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"Director"`
}

type Director struct{
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}

var movies []Movie

func main()  {
	r := mux.NewRouter()

	// seeding movies in
	movies = append(movies, Movie{ID: "1", Isbn:"444447", Title: "Movie one", Director : &Director{firstname: "Mark", lastname: "pro"}})
	movies = append(movies, Movie{ID: "2", Isbn:"444448", Title: "Movie two", Director : &Director{firstname: "Mark", lastname: "pro"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie() {

}

func createMovie() {

}

func updateMovie() {

}

func deleteMovie() {

}