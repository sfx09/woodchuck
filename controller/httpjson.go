package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// Utility functions to write JSON content to HTTP responses

func respondWithJSON(w http.ResponseWriter, status int, msg any) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling JSON %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithErr(w http.ResponseWriter, status int, msg string) {
	type Response struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, status, Response{
		Error: msg,
	})
}
