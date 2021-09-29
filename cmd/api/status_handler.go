package main 

import (
	"net/http"
	"encoding/json"
)

func (app *Application) statusHandler(w http.ResponseWriter, r *http.Request){
	currentStatus := AppStatus {
			Status      : "available",
			Environment : app.config.env,
			Version     : version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(currentStatus)
	if err != nil {
		app.logger.Printf("Error sending the json response %s\n", err)
	}
}