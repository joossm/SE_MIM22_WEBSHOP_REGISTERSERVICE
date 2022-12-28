package handler

import (
	"SE_MIM22_WEBSHOP_MONO/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		responseWriter.Write([]byte("GET"))
	case "POST":
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				log.Fatalf("Error while unmarshalling: %v", jsonErr)
				responseWriter.Write([]byte("{ERROR}"))
			}
			// Find user in database
			// If user is found, return true as json
			// If user is not found, return false as json
			db := openDB()
			result, err := db.Query("SELECT Id, Username, Password FROM users WHERE Username = ? AND Password = ?", user.Username, user.Password)
			errorHandler(err)
			var users []model.User
			for result.Next() {
				var user model.User
				err = result.Scan(&user.Id, &user.Username, &user.Password)

				users = append(users, user)
			}
			defer closeDB(db)
			fmt.Println(user.Username + " " + user.Password)
		}
		responseWriter.Write([]byte("POST"))
	}
}
func Register(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		responseWriter.Write([]byte("GET"))
	case "POST":
		request.Body.Read()
		responseWriter.Write([]byte("POST"))
	}
}
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
func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/todoapi")
	errorHandler(err)
	return db
}
func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
