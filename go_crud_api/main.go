package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/movies" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	// return the existing movies array
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Movie Not Found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/movies" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.IntN(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "PUT" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	http.Error(w, "Movie Not Found", http.StatusNotFound)
}

func main() {

	// setup initial movies data
	// NOTE How append function works in Golang, it takes two arguments, the first is the slice and the second is the element to add to the slice
	//
	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "845690", Title: "Movie Two", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	// Add routes here using gorilla mux
	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start the server with router as the handler
	fmt.Println("Server started on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}

}
