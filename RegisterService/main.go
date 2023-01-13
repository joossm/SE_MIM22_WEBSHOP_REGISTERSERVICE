package main

import (
	"log"
	"net/http"

	"SE_MIM22_WEBSHOP_MONO/handler"
)

func main() { // Server
	var serveMux = http.NewServeMux()
	fileServer := http.FileServer(http.Dir("/view"))
	serveMux.Handle("/", http.StripPrefix("/view", fileServer))
	serveMux.HandleFunc("/login", handler.Login)
	serveMux.HandleFunc("/register", handler.Register)
	serveMux.HandleFunc("/getAllBooks", handler.GetAllBooks)
	serveMux.HandleFunc("/getBookById", handler.GetBookByID)
	serveMux.HandleFunc("/placeOrder", handler.PlaceOrder)
	serveMux.HandleFunc("/getOrdersByUserId", handler.GetOrdersByUserId)
	log.Printf("About to listen on 8443. Go to http://127.0.0.1:8443/register\n Go to http://127.0.0.1:8443/login")
	log.Fatal(http.ListenAndServe(":8443", serveMux))
}