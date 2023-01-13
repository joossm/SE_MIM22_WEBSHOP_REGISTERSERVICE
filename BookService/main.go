package main

import (
	"SE_MIM22_WEBSHOP_MONO/BookService/handler"
	"log"
	"net/http"
	"time"
)

func main() { // Server
	var serveMux = http.NewServeMux()
	fileServer := http.FileServer(http.Dir("/view"))
	serveMux.Handle("/", http.StripPrefix("/view", fileServer))
	serveMux.HandleFunc("/getAllBooks", handler.GetAllBooks)
	serveMux.HandleFunc("/getBookById", handler.GetBookByID)
	log.Printf("\n\n\tBOOKSERVICE\n\nAbout to listen on 8440.\nGet All Books: http://127.0.0.1:8440/getAllBooks\nGet Books By ID: http://127.0.0.1:8440/getBookById?id=1")
	server := &http.Server{
		Addr:              ":8440",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
