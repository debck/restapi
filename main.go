package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/restapi/config"
	. "github.com/restapi/model"
	. "github.com/restapi/controller"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var db = BooksDB{}


func AllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := db.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

// GET a book by its ID
func FindBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := db.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	respondWithJson(w, http.StatusOK, book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	book.ID = bson.NewObjectId()
	if err := db.Insert(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Books
	if err := db.Update(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := db.Delete(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
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