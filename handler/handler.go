package handler

import (
	"SE_MIM22_WEBSHOP_REGISTERSERVICE/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				_, responseErr := responseWriter.Write([]byte("{ERROR}"))
				errorHandler(responseErr)
				return
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
			_, dbErr := db.Query("INSERT INTO users (Username, Password, Firstname, Lastname, HouseNumber, Street, ZipCode, City, Email, Phone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				user.Username, user.Password, user.Firstname, user.Lastname, user.HouseNumber, user.Street, user.ZipCode, user.City, user.Email, user.Phone)
			errorHandler(dbErr)
			_, responseErr := responseWriter.Write([]byte("{true}"))
			errorHandler(responseErr)
			return
		}
	default:
		_, err := responseWriter.Write([]byte("THIS IS A POST REQUEST"))
		errorHandler(err)
		return
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	fmt.Println("Opening DB")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	db.SetMaxIdleConns(0)
	errorHandler(err)
	return db
}
func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
