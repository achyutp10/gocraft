package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mv "smsg/internal/api/middlewares"
	"smsg/internal/api/router"
	"smsg/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("D:/Important/Coding/GoL/prac/z-backend/smsg/cmd/api/.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	_, err = sqlconnect.ConnectDb()

	if err != nil {
		fmt.Println("Error:", err)
	}

	port := os.Getenv("API_PORT")

	cert := "D:/Important/Coding/GoL/prac/z-backend/smsg/cmd/api/cert.pem"
	key := "D:/Important/Coding/GoL/prac/z-backend/smsg/cmd/api/key.pem"

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
	// secureMux := utils.ApplyMiddlewares(mux, mv.Hpp(hppOptions), mv.Compression, mv.SecurityHeaders, mv.ResponseTimeMiddleware, rl.Middleware, mv.Cors)
	router := router.Router()
	secureMux := mv.SecurityHeaders(router)

	// Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler: mv.SecurityHeaders(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
