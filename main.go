package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Init books var as a slice Book struct
var books []Book

func main() {
	// Initialization of Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "LotR", Author: &Author{Firstname: "J.R.R", Lastname: "Tolkien"}})
	books = append(books, Book{ID: "2", Isbn: "845632", Title: "HP", Author: &Author{Firstname: "J.K", Lastname: "Rowling"}})

	// Sets routes
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Adds logger to server
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	// Boots up the server
	log.Println("Server up! Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
