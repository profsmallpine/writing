package main

import (
	"embed"
	"log"
	"os"

	"github.com/profsmallpine/writing/app"
)

// TODO: improve UX/css for reading
// TODO: hide nav on scroll down and show on scroll up

//go:embed tmpl/*.tmpl tmpl/**/*.tmpl
var files embed.FS

func main() {
	// Setup logger.
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	// Setup ranger
	rng, err := app.New(files)
	if err != nil {
		logger.Fatal("could not setup ranger: ", err)
	}

	// Start the web server until receiving a signal to stop
	if err := rng.Guide(); err != nil {
		logger.Fatal("could not guide: ", err)
	}
}
