package main

import (
	"html/template"
	"log"
	"net/http"
)

// Handler is http struct for passing services to the router
type handler struct {
	Logger *log.Logger
}

func (h *handler) goHome(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/index.tmpl", nil)
}

func respond(logger *log.Logger, w http.ResponseWriter, r *http.Request, layout string, data interface{}) {
	// Parse static files.
	tmpl := template.Must(template.New("base.tmpl").ParseFiles(
		"./tmpl/base.tmpl",
		layout,
	))
	err := tmpl.Funcs(template.FuncMap{}).Execute(w, data)

	// Log error is compiling template fails.
	if err != nil {
		logger.Println("Template execution error: ", err.Error(), layout)
		return
	}
}
