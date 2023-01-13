package handler

import (
	"encoding/json"
	"net/http"

	"SE_MIM22_WEBSHOP_MONO/model"
)

func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := model.OpenDB()
		defer model.CloseDB(db)
		result, err := db.Query("SELECT * FROM books")
		errorHandler(err)
		var books []model.Book
		for result.Next() {
			var book model.Book
			err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
			errorHandler(err)
			books = append(books, book)
		}
		json, err := json.Marshal(books)
		errorHandler(err)
		responseWriter.Write(json)
	default:
		responseWriter.Write([]byte("THIS IS A GET REQUEST"))
	}
}

func GetBookByID(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := model.OpenDB()
		defer model.CloseDB(db)
		result, err := db.Query("SELECT * FROM books WHERE Id = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var books []model.Book
		for result.Next() {
			var book model.Book
			err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
			errorHandler(err)
			books = append(books, book)
		}
		json, err := json.Marshal(books)
		errorHandler(err)
		responseWriter.Write(json)
	default:
		responseWriter.Write([]byte("THIS IS A GET REQUEST"))
	}
}

func errorHandler(err error) {
	if err != nil {
		print(err)
	}
}
