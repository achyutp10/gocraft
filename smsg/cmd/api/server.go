package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
	// teachers/{id}
	// teachers/9
	// teachers/? &
	// teachers/?key=value&query=value2&sortby=email&sortorder=ASC

	switch r.Method {
	case http.MethodGet:

		// Path Params
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		fmt.Println("Path:", path)
		userID := strings.TrimSuffix(path, "/")
		fmt.Println("The ID is: ", userID)

		// Query params
		fmt.Println("Query params:", r.URL.Query())
		querParams := r.URL.Query()
		sortby := querParams.Get("sortby")
		key := querParams.Get("key")
		sortorder := querParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "DESC"
		}

		fmt.Printf("Sort by: %v, Sort order: %v, key: %v", sortby, sortorder, key)

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

	// fmt.Fprintf(w, "Hello Root Route")
	// w.Write([]byte("Hello Teachers Route"))
	// fmt.Println("Hello Teachers route")
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

	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello students Route"))
	fmt.Println("Hello students route")
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

	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello execs Route"))
	fmt.Println("Hello execs route")
}

func main() {

	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
