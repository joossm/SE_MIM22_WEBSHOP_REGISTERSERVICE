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
		fmt.Println("Register was executed")
		if request.Body != nil {
			fmt.Println("Body not nil")
			body, _ := io.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				js, err := json.Marshal("Error")
				errorHandler(err)
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
			fmt.Println("No json error")
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Username FROM users WHERE Username = ?", user.Username)
			fmt.Println("result: ", result)
			errorHandler(err)
			fmt.Println("Query executed")
			var users []model.User
			if result.Next() == true {
				for result.Next() {
					var user model.User
					err = result.Scan(&user.Id, &user.Username, &user.Password)
					fmt.Println("user: ", user.Username, user.Password)
					users = append(users, user)
				}
				if users != nil {
					js, err := json.Marshal("already exists")
					errorHandler(err)
					_, responseErr := responseWriter.Write(js)
					errorHandler(responseErr)
					return
				}
			} else {
				// GET MAX ID
				result, err := db.Query("SELECT MAX(Id) FROM users")
				errorHandler(err)
				var maxId int
				if result != nil {
					for result.Next() {
						err = result.Scan(&maxId)
						errorHandler(err)
					}
				}
				maxId++
				fmt.Println("result is nil | execute insert")
				res, err := db.Query("INSERT INTO users (Id, Username, Password, Firstname, Lastname, HouseNumber, Street, ZipCode, City, Email, Phone) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
					maxId, user.Username, user.Password, user.Firstname, user.Lastname, user.HouseNumber, user.Street, user.ZipCode, user.City, user.Email, user.Phone)
				fmt.Println(res)
				errorHandler(err)
				js, err := json.Marshal("true")
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
		}
	default:
		js, err := json.Marshal("THIS IS A POST REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func Error(responseWriter http.ResponseWriter, request *http.Request) {
	// This is just a test function to create an error
	Error(responseWriter, request)
	panic("ERROR")
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
