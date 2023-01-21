package main

import (
	"SE_MIM22_WEBSHOP_REGISTERSERVICE/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/register", handler.Register)
	serveMux.HandleFunc("/error", handler.Error)
	log.Printf("\n\n\tREGISTERSERVICE\n\nAbout to listen on Port: 8442.\n\nSUPPORTED REQUESTS:\n" +
		"GET:\n" +
		"Create Error on: http://127.0.0.1:8442/error\n" +
		"POST:\n" +
		"Register on: http://127.0.0.1:8442/register requires a JSON Body with the following format:\n" +
		"{\n    \"Username\": \"mmuster\",\n    \"Password\": \"password\",\n    \"Firstname\": \"Max\",\n   " +
		" \"Lastname\": \"Muster\",\n    \"Housenumber\": \"1\",\n    \"Street\": \"Musterstr.\",\n  " +
		"  \"Zipcode\": \"01234\",\n    \"City\": \"Musterstadt\",\n    \"Email\": \"max.muster@mail.com\",\n  " +
		"  \"Phone\": \"012345678910\"\n  }")
	server := &http.Server{
		Addr:              ":8442",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
