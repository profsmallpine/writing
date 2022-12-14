package main

import (
	"log"
	"os"

	"github.com/profsmallpine/mid/app"
)

// TODO: improve UX/css for reading
// TODO: hide nav on scroll down and show on scroll up

func main() {
	// Setup logger.
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	// Setup ranger
	rng, err := app.New()
	if err != nil {
		logger.Fatal("could not setup ranger: ", err)
	}

	// Start the web server until receiving a signal to stop
	if err := rng.Guide(); err != nil {
		logger.Fatal("could not guide: ", err)
	}
}
