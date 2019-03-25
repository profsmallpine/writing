package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/profsmallpine/mid/postgres"
)

// TODO: improve UX/css for reading
// TODO: hide nav on scroll down and show on scroll up

// NOTE: move writing to content from database and single handler, blocking:
// 			 schema based on writing organization

func main() {
	// Setup logger.
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	// Load .env file.
	if err := godotenv.Load(); err != nil {
		panic("could not load env!")
	}

	// Connect/migrate database.
	db, err := postgres.Setup(&postgres.DBConfig{
		Host:       os.Getenv("PG_HOST"),
		Port:       os.Getenv("PG_PORT"),
		Name:       os.Getenv("PG_NAME"),
		User:       os.Getenv("PG_USER"),
		Password:   os.Getenv("PG_PASSWORD"),
		VerboseLog: false,
	})
	if err != nil {
		logger.Panicf("Failed to set up database: %s", err)
		return
	}

	// Build handler.
	h := handler{Logger: logger, DB: db}

	// Setup routes.
	router := buildRoutes(h)

	// Run server.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic("could not start server!")
	}
}
