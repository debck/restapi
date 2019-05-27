package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented")
}

func FindBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", AllBooks).Methods("GET")
	r.HandleFunc("/books", AddBook).Methods("POST")
	r.HandleFunc("/books", UpdateBook).Methods("PUT")
	r.HandleFunc("/books", DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}", FindBook).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}