package handler

import (
	"SE_MIM22_WEBSHOP_MONO/model"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	case "POST":
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				log.Fatalf("Error while unmarshalling: %v", jsonErr)
				responseWriter.Write([]byte("{ERROR}"))
			}
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Id, Username, Password FROM users WHERE Username = ? AND Password = ?", user.Username, user.Password)
			errorHandler(err)
			var users []model.User
			for result.Next() {
				var user model.User
				err = result.Scan(&user.Id, &user.Username, &user.Password)

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
	case "POST":
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				log.Fatalf("Error while unmarshalling: %v", jsonErr)
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
			responseWriter.Write([]byte("{false}"))
			return
		}
	}
}

/*
	func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			responseWriter.Write([]byte("GET"))
		case "POST":
			request.Body.Read()
			responseWriter.Write([]byte("POST"))
		}
	}

	func GetBook(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			responseWriter.Write([]byte("GET"))
		case "POST":
			request.Body.Read()
			responseWriter.Write([]byte("POST"))
		}
	}

	func AddToBasket(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			responseWriter.Write([]byte("GET"))
		case "POST":
			request.Body.Read()
			responseWriter.Write([]byte("POST"))
		}
	}
*/
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
		log.Fatal(err)
	}
}
