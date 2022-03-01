package main

import (
	"github.com/gorilla/mux"
	"github.com/jacoblbeck/fibonacci-api/api"
)

//router sets the api endpoints for the application and docker registry.
func router(options ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()

	router.Use(options...)

	router.HandleFunc("/api/v1/fibonacci/next", api.Next).Methods("GET")
	router.HandleFunc("/api/v1/fibonacci/current", api.Current).Methods("GET")
	router.HandleFunc("/api/v1/fibonacci/previous", api.Previous).Methods("GET")

	router.HandleFunc("/health", api.Health).Methods("GET")

	return router

}
