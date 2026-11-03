package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mv "smsg/internal/api/middlewares"
	"smsg/internal/api/router"
)

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

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
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
