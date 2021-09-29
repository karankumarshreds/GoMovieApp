package main 

import (
	"os"
	"fmt"
	"log"
	"flag"
	"net/http"
	"encoding/json"
	_"github.com/gorilla/mux"
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
}

func main() {
	var app Application

	flag.IntVar(&app.config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&app.config.env, "env", "development", "Application environment (development|production)")
	flag.Parse()
	log.Printf("Running the application on port %v in %v mode...", app.config.port, app.config.env)

	logger := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	app.logger = logger 

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request){
		
		currentStatus := AppStatus {
			Status      : "available",
			Environment : app.config.env,
			Version     : version,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(currentStatus)
		if err != nil {
			log.Printf("Error sending the json response %s\n", err)
		}

	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", app.config.port), nil)
	if err != nil {
		log.Fatalf("Unable to start the application")
	}
}