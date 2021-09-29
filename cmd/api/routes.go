package main 

import "github.com/gorilla/mux"

func (app *Application) routes() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/status", app.statusHandler).Methods("GET")
	router.HandleFunc("/api/movie/{id}", app.getOneMovie).Methods("GET")
	router.HandleFunc("/api/movie/all", app.getAllMovies).Methods("GET")
	
	return router
}