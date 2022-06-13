package main

import (
	// "fmt"	// FOR FORMATTING I/O
	"encoding/json" // FOR WORKING WITH JSON
	"log"           // TO LOG ERRORS
	"net/http"      // TO WORK WITH HTTP, WE CREATE APIs WITH THIS

	"math/rand" // FOR RANDOM NUMBER GENERATION
	"strconv"   // FOR CONVERSION TO STRING

	"github.com/go-chi/chi" // ROUTER FRAMEWORK CHI
)

// BOOK STRUCT (MODEL)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// AUTHOR STRUCT
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// INIT BOOKS VAR AS A SLICE BOOK STRUCT
var books []Book

// GET ALL BOOKS - FUNCTION
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET ONE BOOK - FUNCTION
func getOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.Vars(r) // GET PARAMS
	// LOOP THROUGH BOOKS AND FIND WITH ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// CREATE A NEW BOOK - FUNTION
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // MOCK ID - NOT SAFE, MAY REPEAT
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// UPDATE A BOOK - FUNCTION
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.Vars(r) // GET PARAMS
	// LOOP THROUGH BOOKS AND FIND WITH ID
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(books)
}

// DELETE A BOOK - FUNCTION
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.Vars(r) // GET PARAMS
	// LOOP THROUGH BOOKS AND FIND WITH ID
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	// INIT ROUTER
	router := chi.NewRouter()

	// MOCK DATA - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "562365", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "289041", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})

	// ROUTE HANDLERS / ENDPOINTS
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getOneBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// TO RUN SERVER
	// WE USE HTTP PACKAGE -> LISTEN AND SERVE METHOD
	// TAKES PORTS AND ROUTER
	log.Fatal(http.ListenAndServe(":8000", router))

}
