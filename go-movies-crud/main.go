package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

type CustomError struct {
	Message string `json:"message"`
}

var movies []Movie

func sendErrorMsg(w http.ResponseWriter, r *http.Request) {
	errMsg := CustomError{Message: "Movie not found"}

	jsonData, jsonErr := json.Marshal(errMsg)

	if jsonErr != nil {
		fmt.Println("Error encoding JSON:", jsonErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, writeErr := w.Write(jsonData)

	if writeErr != nil {
		fmt.Println("Error writing JSON reponse", writeErr)
	}
}

func getMovieById(id string) (int, Movie) {
	var idx int = -1
	var itm Movie
	for index, item := range movies {
		if item.ID == id {
			idx = index
			itm = item
		}
	}

	return idx, itm
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	movieId := params["id"]

	index, _ := getMovieById(movieId)
	if index > -1 {
		movies = append(movies[:index], movies[index+1:]...)
		json.NewEncoder(w).Encode(movies)
		return
	}

	sendErrorMsg(w, r)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that we're returning JSON.

	w.Header().Set("Content-Type", "application.json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	sendErrorMsg(w, r)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// validitions are missing

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	movieId := params["id"]
	index, _ := getMovieById(movieId)
	var movie Movie

	if index > -1 {
		movies = append(movies[:index], movies[index+1:]...)
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = movieId
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
		return
	}

	sendErrorMsg(w, r)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438278", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")

	// when deploying to production remove localhost prefix
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
