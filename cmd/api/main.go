package main 

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"encoding/json"
)

const version = "1.0.0"

type AppStatus struct {
	Status       string    `json:"status"`
	Environment  string    `json:"environment"` 
	Version      string    `json:"version"` 
}

type config struct {
	port int 
	env string 
}

func main() {
	var cfg config 

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()
	log.Printf("Running the application on port %v in %v mode...", cfg.port, cfg.env)
	

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request){
		
		currentStatus := AppStatus {
			Status      : "available",
			Environment : cfg.env,
			Version     : version,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(currentStatus)
		if err !=nil {
			log.Printf("Error sending the json response %s\n", err)
		}

	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), nil)
	if err !=nil {
		log.Fatalf("Unable to start the application")
	}
}