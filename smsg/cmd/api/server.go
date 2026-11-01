package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mv "smsg/internal/api/middlewares"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		w.Write([]byte("Hello GET method on teachers route"))
		// fmt.Println("Hello GET method on teachers route")

	case http.MethodPost:
		w.Write([]byte("Hello POST method on teachers route"))
		fmt.Println("Hello POST method on teachers route")

	case http.MethodPut:
		w.Write([]byte("Hello Put method on teachers route"))
		fmt.Println("Hello Put method on teachers route")

	case http.MethodPatch:
		w.Write([]byte("Hello Patch method on teachers route"))
		fmt.Println("Hello Patch method on teachers route")

	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on teachers route"))
		fmt.Println("Hello Delete method on teachers route")
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on students route"))
		fmt.Println("Hello GET method on students route")

	case http.MethodPost:
		w.Write([]byte("Hello POST method on students route"))
		fmt.Println("Hello POST method on students route")

	case http.MethodPut:
		w.Write([]byte("Hello Put method on students route"))
		fmt.Println("Hello Put method on students route")

	case http.MethodPatch:
		w.Write([]byte("Hello Patch method on students route"))
		fmt.Println("Hello Patch method on students route")

	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on students route"))
		fmt.Println("Hello Delete method on students route")
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on execs route"))
		fmt.Println("Hello GET method on execs route")

	case http.MethodPost:
		w.Write([]byte("Hello POST method on execs route"))
		fmt.Println("Hello POST method on execs route")

	case http.MethodPut:
		w.Write([]byte("Hello Put method on execs route"))
		fmt.Println("Hello Put method on execs route")

	case http.MethodPatch:
		w.Write([]byte("Hello Patch method on execs route"))
		fmt.Println("Hello Patch method on execs route")

	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on execs route"))
		fmt.Println("Hello Delete method on execs route")
	}

}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		Handler: mv.SecurityHeaders(mux),
		// Handler:   middlewares.Cors(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
