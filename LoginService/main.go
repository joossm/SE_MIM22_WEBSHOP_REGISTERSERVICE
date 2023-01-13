package main

import (
	"SE_MIM22_WEBSHOP_MONO/LoginService/handler"
	"log"
	"net/http"
	"time"
)

func main() { // Server
	var serveMux = http.NewServeMux()
	fileServer := http.FileServer(http.Dir("/view"))
	serveMux.Handle("/", http.StripPrefix("/view", fileServer))
	serveMux.HandleFunc("/login", handler.Login)
	log.Printf("\n\n\tLOGINSERVICE\n\nAbout to listen on 8441.\nGet All Books: http://127.0.0.1:8441/login")
	server := &http.Server{
		Addr:              ":8441",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
