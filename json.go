package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResp struct {
	Error string `json:"error"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	// marshal the payload struct into a []byte
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Write(resp)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5xx error: %s", msg)
	}

	w.WriteHeader(code)
	w.Header().Set("Conntent-Type", "application/json")

	// create a []byte from the msg

	errorRespBody := errorResp{
		Error: msg,
	}

	resp, err := json.Marshal(errorRespBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Write(resp)

}
