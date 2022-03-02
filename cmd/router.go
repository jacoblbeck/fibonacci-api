package main

import (
	"github.com/gorilla/mux"
	"github.com/jacoblbeck/fibonacci-api/api"
)

//router sets the api endpoints for the application.
func router(options ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()

	router.Use(options...)

	router.HandleFunc("/api/v1/fibonacci/next", api.GetNext).Methods("GET")
	router.HandleFunc("/api/v1/fibonacci/current", api.GetCurrent).Methods("GET")
	router.HandleFunc("/api/v1/fibonacci/previous", api.GetPrevious).Methods("GET")

	router.HandleFunc("/api/v1/fibonacci/reset", api.ResetSequence).Methods("GET")

	router.HandleFunc("/health", api.Health).Methods("GET")

	return router

}
