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
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var me []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content Type", "Application/json")
	json.NewEncoder(w).Encode(me)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content type", "Application/json")
	// get the id from the user
	params := mux.Vars(r)

	for i,val := range me{
		if val.ID == params["id"]{
			// using the slice to delete that id 
			me = append(me[:i], me[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(me)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the id from the user
	params := mux.Vars(r)

	for _,val := range me{
		if val.ID == params["id"]{
			// sending the particular movie
			json.NewEncoder(w).Encode(val)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "application/json")
	//  define a variable called
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(10000000))
	me = append(me, newMovie)

	json.NewEncoder(w).Encode(newMovie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var newMovie Movie

	for i,val := range me {
		if val.ID == params["id"] {
			me = append(me[:i],me[i+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID = params["id"]
			me = append(me, newMovie)
			json.NewEncoder(w).Encode(newMovie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	// slices
	me = append(me, Movie{
		ID: "1",
		Isbn: "43510",
		Title: "Movie one",
		Director: &Director{
			Firstname: "John",
			Lastname: "Doe",
		},
	})

	me = append(me, Movie{
		ID: "2",
		Isbn: "78110",
		Title: "Movie two",
		Director: &Director{
			Firstname: "Mark",
			Lastname: "Andrew",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server at port: 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}