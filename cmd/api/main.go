package main 

import (
	"os"
	"fmt"
	"log"
	"flag"
	"time"
	"net/http"
	_"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoMovieApp/pkg/utils"
)

const version = "1.0.0"

type AppStatus struct {
	Status       string    `json:"status"`
	Environment  string    `json:"environment"` 
	Version      string    `json:"version"` 
}


type Application struct {
	config struct {
		port int 
		env string 
	}
	logger *log.Logger
	utils utils.Utils
}

func main() {
	var app Application

	flag.IntVar(&app.config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&app.config.env, "env", "development", "Application environment (development|production)")
	flag.Parse()
	

	logger := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	app.logger = logger 

	srv := &http.Server{
		Addr         : fmt.Sprintf(":%v", app.config.port),
		Handler      : app.routes(),
		IdleTimeout  : time.Minute,
		ReadTimeout  : 10 * time.Second,
		WriteTimeout : 30 * time.Second,
	}
	logger.Printf("Running the application on port %v in %v mode...\n", app.config.port, app.config.env)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start the application")
	}
}