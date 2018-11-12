package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Setup logger.
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	// Load .env file.
	if err := godotenv.Load(); err != nil {
		logger.Println("error loading .env file")
		return
	}

	// Build handler.
	// TODO: attach dependencies for shared functionality.
	h := handler{Logger: logger}

	// Setup routes.
	router := buildRoutes(h)

	// Run server.
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
