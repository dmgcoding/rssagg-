package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Printf("responding with 5xx error: %v",msg)
	}

	type errorResponse struct{
		Error string `json:"error"`
	}
	respondWithJSON(w,code,errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat,err := json.Marshal(payload)
	if err != nil{
		log.Printf("Failed to marshal json response: $v",payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}