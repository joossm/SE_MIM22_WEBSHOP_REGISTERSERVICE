package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"SE_MIM22_WEBSHOP_MONO/model"
	_ "github.com/go-sql-driver/mysql"
)

const post = "POST"

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	case post:
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
				return
			}
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Id, Username, Password FROM users WHERE Username = ? AND Password = ?", user.Username, user.Password)
			errorHandler(err)
			var users []model.User
			for result.Next() {
				var user model.User
				err = result.Scan(&user.Id, &user.Username, &user.Password)
				errorHandler(err)
				users = append(users, user)
			}
			for _, iUser := range users {
				fmt.Println(user.Username + " " + user.Password)
				fmt.Println(iUser.Username + " " + iUser.Password)
				if iUser.Username == user.Username && iUser.Password == user.Password {
					responseWriter.Write([]byte("{true}"))
					return
				}
			}
			responseWriter.Write([]byte("{false}"))
			return
		}
		responseWriter.Write([]byte("{false}"))
	}
}

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	case post:
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
			}
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Username FROM users WHERE Username = ?", user.Username)
			errorHandler(err)
			var users []model.User
			for result.Next() {
				var user model.User
				err = result.Scan(&user.Id, &user.Username, &user.Password)
				users = append(users, user)
			}

			if users != nil {
				responseWriter.Write([]byte("{already exists}"))
				return
			} else {
				db.Query("INSERT INTO users (Username, Password, Firstname, Lastname, HouseNumber, Street, ZipCode, City, Email, Phone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
					user.Username, user.Password, user.Firstname, user.Lastname, user.HouseNumber, user.Street, user.ZipCode, user.City, user.Email, user.Phone)
				responseWriter.Write([]byte("{true}"))
				return
			}
		}
	}
}

func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
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
		db := openDB()
		defer closeDB(db)
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

func PlaceOrder(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case post:
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			order := model.Order{}
			jsonErr := json.Unmarshal(body, &order)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
				return
			}
			db := openDB()
			defer closeDB(db)
			db.Query("INSERT INTO orders (produktId, userId, Amount) VALUES (?, ?, ?)",
				order.ProduktId, order.UserId, order.Amount)
			responseWriter.Write([]byte("{true}"))
			return
		}
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	}
}
func GetOrdersByUserId(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT produktId,userId, amount FROM orders WHERE userId = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var orders []model.Order
		for result.Next() {
			var order model.Order
			err = result.Scan(&order.ProduktId, &order.UserId, &order.Amount)
			errorHandler(err)
			orders = append(orders, order)
		}
		json, err := json.Marshal(orders)
		errorHandler(err)
		responseWriter.Write(json)
	default:
		responseWriter.Write([]byte("THIS IS A GET REQUEST"))
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/books")
	errorHandler(err)
	return db
}
func errorHandler(err error) {
	if err != nil {
		print(err)
	}
}