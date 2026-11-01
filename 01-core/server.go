package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server!")
	})

	const port string = ":8080"

	fmt.Println("Server Listening on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}

}

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main2() {

	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello root route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method on teachers route"))
			fmt.Println("Hello GET method on teachers route")
		case http.MethodPost:

			// Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}
			fmt.Println("Form:", r.Form)

			// Prepare response data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0]
			}
			fmt.Println("Processed Response Map:", response)

			// RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()
			fmt.Println("RAW Body:", string(body))

			// IF you expect json data, then unmarshal it
			var userInstance user
			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled json into an instance of user struct:", userInstance)
			fmt.Println("Name:", userInstance.Name)
			fmt.Println("Age:", userInstance.Age)

			// Prepare response data
			response1 := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0]
			}

			err = json.Unmarshal(body, &response1)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled json into a map", response1)

			// Access the request details
			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("ContentLength:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Addr:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("URL:", r.URL)
			fmt.Println("UserAgent:", r.UserAgent())
			fmt.Println("Port:", r.URL.Port())
			fmt.Println("URL Scheme:", r.URL.Scheme)

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
		w.Write([]byte("Hello Teachers Route"))
		fmt.Println("Hello Teachers route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello students Route"))
		fmt.Println("Hello students route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello execs Route"))
		fmt.Println("Hello execs route")
	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}


// ==================================

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
