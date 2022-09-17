package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

// Init books var as slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get any params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get any params
	for index, item := range books {
		if item.ID == params["id"] {
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			books[index] = book
			json.NewEncoder(w).Encode(books)
		}
	}
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get any params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book 1", Author: &Author{
		Firstname: "John",
		Lastname:  "Doe",
	}})
	books = append(books, Book{ID: "2", Isbn: "43438743", Title: "Book 2", Author: &Author{
		Firstname: "Ali",
		Lastname:  "Daei",
	}})
	books = append(books, Book{ID: "3", Isbn: "43243", Title: "Book 3", Author: &Author{
		Firstname: "Reza",
		Lastname:  "Alizadeh",
	}})
	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
