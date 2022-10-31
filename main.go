package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	M "example.com/crud/models"
	L "example.com/crud/logic"
)

var movies []M.Movie

func main(){
	r := mux.NewRouter()

	movies = append(movies, M.Movie{ID:"1", Isbn:"24601", Title:"Movie 1", Director: &M.Director{Firstname:"John", Lastname:"Doe"}})
	movies = append(movies, M.Movie{ID:"2", Isbn:"24602", Title:"Movie 2", Director: &M.Director{Firstname:"Jane", Lastname:"Smith"}})

	r.HandleFunc("/movie", L.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", L.GetMovie).Methods("GET")
	r.HandleFunc("/movie", L.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", L.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", L.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}