package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// TODO: Add transition to section from navbar
// TODO: make ^ work from writing pages
// TODO: write DS blog and add to website

// NOTE: move writing to content from database and single handler, blocking:
// 			 schema based on writing organization

func main() {
	// Setup logger.
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	// Load .env file.
	if err := godotenv.Load(); err != nil {
		logger.Println("error loading .env file")
		return
	}

	// Build handler.
	h := handler{Logger: logger}

	// Setup routes.
	router := buildRoutes(h)

	// Run server.
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
