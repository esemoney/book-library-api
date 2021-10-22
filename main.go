package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"
	// "time"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Homepage Endpoint Hit 😊")

}

type BookInfo struct{
	BookID string `json:"ID"`
	BookName string `json:"name"`
	BookAuthor string `json:"author"`
	TimeAdded string `json:"added_at"`
}

// i defined a slice to store the  books since i wouldn't be using a database. slice instead of array because slices can be resized dynamically

var bookCollection []BookInfo

// i also create and initialise a variable to generate and assign the id for each book i add to the library. to add a new book, i just need to increaase this and assign it to the BookID field of the BookInfo struct
var prevBookID int = 0

// next, i create the api for adding a book to the library
func addBook(w http.ResponseWriter, r *http.Request) {

	var book BookInfo
	json.NewDecoder(r.Body).Decode(&book)
	prevBookID ++
	book.BookID = strconv.Itoa(prevBookID)
	bookCollection = append(bookCollection, book)
	// below sets that the api returns a json content as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	fmt.Fprintf(w, "Test POST ENDPOINT WORKED")


}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookCollection)
	fmt.Println("Endpoint Hit: All books endpoint")

}

// for my delete function, first, the function gets the api from the request, then it iterates through the bookCollection slice and finds the book with the input id. once it finds the book, it removes it from the slice using the append function. 

// interesting how go uses append to delete stuff


func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	inputBookID := params["ID"]
	for i, book := range bookCollection {
		if book.BookID== inputBookID {
			bookCollection = append(bookCollection[:i], bookCollection[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

}


func main() {

	router := mux.NewRouter()
	// to display the homepage endpoint
	router.HandleFunc("/", homePage)

	// post request to add the book to library
	router.HandleFunc("/book", addBook).Methods("POST")

	// get request to return json array of all the books in library
	router.HandleFunc("/books", getBooks).Methods("GET")

	// delete request to delete the book with provided id
	router.HandleFunc("/delete/{id}", deleteBook).Methods("DELETE")

	// to run it!
	log.Fatal(http.ListenAndServe(":8081", router))

}


