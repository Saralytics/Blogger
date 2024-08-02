package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	type resp struct {
		Response string `json:"resp"`
	}
	respondWithJSON(w, http.StatusOK, resp{
		Response: "System ready",
	})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {

	respondWithError(w, http.StatusInternalServerError, "System not ready")
}
