package main 

import (
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoMovieApp/pkg/models"
)

func (app *Application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.logger.Println("Error", err)
	}
	app.logger.Println("ID is", id)
	movie := models.Movie{
		CreatedAt: time.Now(),
		Description: "Some random description",
		Title: "Batman",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (app *Application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	
}