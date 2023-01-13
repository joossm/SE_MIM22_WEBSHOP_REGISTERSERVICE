package main

import (
	"SE_MIM22_WEBSHOP_MONO/RegisterService/handler"
	"log"
	"net/http"
	"time"
)

func main() { // Server
	var serveMux = http.NewServeMux()
	fileServer := http.FileServer(http.Dir("/view"))
	serveMux.Handle("/", http.StripPrefix("/view", fileServer))
	serveMux.HandleFunc("/register", handler.Register)
	log.Printf("\n\n\tREGISTERSERVICE\n\nAbout to listen on 8442.\nGet All Books: http://127.0.0.1:8442/register")
	server := &http.Server{
		Addr:              ":8442",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
