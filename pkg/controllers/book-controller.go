package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/mahdi-eth/go-bookstore/pkg/utils"
	"github.com/mahdi-eth/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(wr http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func GetBookById(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while fetching")
	}
	bookDetails, _ := models.GetBookById(Id)

	res, _ := json.Marshal(bookDetails)
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func CreateBook(wr http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	bookDetails := CreateBook.CreateBook()

	res, _ := json.Marshal(bookDetails)
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func DeleteBook(wr http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while fetching")
	}
	bookDetails := models.DeleteBook(Id)

	res, _ := json.Marshal(bookDetails)
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func UpdateBook(wr http.ResponseWriter, req *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while fetching")
	}
	bookDetails, db := models.GetBookById(Id)

	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}

	db.Save(bookDetails)

	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}