package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func writeResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	writeResponse(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	writeResponse(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	book := createBook.CreateBook()
	res, _ := json.Marshal(book)
	writeResponse(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	writeResponse(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	res, _ := json.Marshal(book)
	writeResponse(w, res)
}
