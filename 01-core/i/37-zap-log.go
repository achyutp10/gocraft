package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in initializing zap logger")
	}
	defer logger.Sync()
	logger.Info("This is an info message")

	logger.Info("User Logged in", zap.String("username", "John doe"), zap.String("method", "GET"))

}
