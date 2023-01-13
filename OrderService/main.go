package main

import (
	"SE_MIM22_WEBSHOP_MONO/OrderService/handler"
	"log"
	"net/http"
	"time"
)

func main() { // Server
	var serveMux = http.NewServeMux()
	fileServer := http.FileServer(http.Dir("/view"))
	serveMux.Handle("/", http.StripPrefix("/view", fileServer))
	serveMux.HandleFunc("/placeOrder", handler.PlaceOrder)
	serveMux.HandleFunc("/getOrdersByUserId", handler.GetOrdersByUserId)
	log.Printf("\n\n\tREGISTERSERVICE\n\nAbout to listen on 8443.\nGet All Books: http://127.0.0.1:8443/placeOrder\nGet Books By ID: http://127.0.0.1:8443/getOrdersByUserId?id=1")
	server := &http.Server{
		Addr:              ":8443",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
