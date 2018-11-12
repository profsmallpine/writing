package main

import (
	"log"
	"net/http"
)

// Handler is http struct for passing services to the router
type handler struct {
	Logger *log.Logger
}

func (h *handler) goHome(w http.ResponseWriter, r *http.Request) {
}
