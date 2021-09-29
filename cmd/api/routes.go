package main 

import "github.com/gorilla/mux"

func (app *Application) routes() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/status", app.statusHandler).Methods("GET")
	
	return router
}