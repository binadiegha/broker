package main

import (
	"encoding/json"
	"net/http"
)


type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"` //for backward compatibility sake of this code 
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "You have reached the broker",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}