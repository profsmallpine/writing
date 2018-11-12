package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type route struct {
	path    string
	method  string
	handler http.HandlerFunc
}

func buildRoutes(h handler) *httprouter.Router {
	// Setup middlewares, adding to this slice will run code before each web
	// request.
	middlewares := []adapter{
		logRequest(h.Logger),
	}

	// Routes setup + middlewares injection.
	router := httprouter.New()
	routes := []route{
		route{path: "/", method: http.MethodGet, handler: h.goHome},
	}
	for _, r := range routes {
		router.Handler(r.method, r.path, chain(
			r.handler,
			middlewares...,
		))
	}
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	return router
}
