package utils 

import (
	"net/http"
	"encoding/json"
)

type Utils struct {}

func (u Utils) WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	// wrapper := make(map[string]interface{})
	// wrapper[wrap] = data
	js, err := json.Marshal(data)
	if err != nil {
		return err 
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}