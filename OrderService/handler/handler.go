package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"SE_MIM22_WEBSHOP_MONO/model"
	_ "github.com/go-sql-driver/mysql"
)

const post = "POST"

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
