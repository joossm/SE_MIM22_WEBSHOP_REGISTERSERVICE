package handler

import (
	"SE_MIM22_WEBSHOP_MONO/model"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const post = "POST"

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		_, err := responseWriter.Write([]byte("THIS IS A POST REQUEST"))
		errorHandler(err)
	case post:
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
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
				errScan := result.Scan(&user.Id, &user.Username, &user.Password)
				errorHandler(errScan)
				users = append(users, user)
			}

			if users != nil {
				_, errResponse := responseWriter.Write([]byte("{already exists}"))
				errorHandler(errResponse)
				return
			}
			db.Query("INSERT INTO users (Username, Password, Firstname, Lastname, HouseNumber, Street, ZipCode, City, Email, Phone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				user.Username, user.Password, user.Firstname, user.Lastname, user.HouseNumber, user.Street, user.ZipCode, user.City, user.Email, user.Phone)
			responseWriter.Write([]byte("{true}"))
			return
		}
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
