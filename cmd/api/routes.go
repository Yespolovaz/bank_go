package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler { 
	router := mux.NewRouter()


	router.HandleFunc("/api/v1/healthcheck", app.healthcheckHandler).Methods("GET")
	
	router.HandleFunc("/api/v1/users", listUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/users/{id}", updateUser).Methods("PUT")

	return app.authenticate(r)
}
