package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mv "smsg/internal/api/middlewares"
	"strconv"
	"strings"
)

// path /teachers/id
// path /teachers/
type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Class     string
	Subject   string
}

var (
	teachers = make(map[int]Teacher)
	// mutex = &sync.Mutex{}
	nextID = 1
)

// Initialize some dummy data

func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Achyut",
		LastName:  "Parajuli",
		Class:     "9A",
		Subject:   "DB",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Adv-DB",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Math",
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}
		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teachers),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// handle path parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found:", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		// call get method handler function
		getTeachersHandler(w, r)

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

	// rl := mv.NewRateLimiter(5, time.Minute)

	// hppOptions := mv.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mv.Cors(rl.Middleware(mv.ResponseTimeMiddleware(mv.Compression(mv.SecurityHeaders(mv.Hpp(hppOptions)(mux))))))
	// secureMux := ApplyMiddlewares(mux, mv.Hpp(hppOptions), mv.Compression, mv.SecurityHeaders, mv.ResponseTimeMiddleware, rl.Middleware, mv.Cors)
	secureMux := mv.SecurityHeaders(mux)

	// Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler: mv.SecurityHeaders(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}

// Middleware is a function that wraps an http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
