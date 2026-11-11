package main

import (
	"log"
	"net/http"

	"cd/internal/config"
	"cd/internal/database"
	"cd/internal/modules/notes"
)

func main() {

	// 1. Load configuration
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// 2. Connect to PostgreSQL
	db, err := database.ConnectPostgres(cfg)

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	defer db.Close()

	log.Println("PostgreSQL connected successfully")

	// 3. Initialize repository
	noteRepository := notes.NewPostgresRepository(db)

	// 4. Initialize service
	noteService := notes.NewService(noteRepository)

	// 5. Initialize handler
	noteHandler := notes.NewHandler(noteService)

	// 6. Initialize router
	mux := http.NewServeMux()

	notes.RegisterRoutes(mux, noteHandler)

	// 7. Start server
	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	log.Println("Server running on port", cfg.AppPort)

	err = server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal("Server failed:", err)
	}
}
