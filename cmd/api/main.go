package main 

import (
	"flag"
	"log"
)

const version = "1.0.0"

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
	
}