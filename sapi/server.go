package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handeling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handeling users")
	})

	port := 3000

	// openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pen -out cert.pen -days 365
	// Load the TLS cert and key
	cert := "cert.pen"
	key := "key.pen"

	// configure TLS

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}

	// http 1.1 server without tls
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start server", err)
	// }

}
