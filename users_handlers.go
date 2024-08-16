package main

import (
	"Blogger/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleUsersCreate(w http.ResponseWriter, r *http.Request) {

	// struct to store the parameters parsed from request
	type parameters struct {
		Name string
	}
	// parse the request
	decoder := json.NewDecoder(r.Body)
	params := parameters{}         // instantiate an instance of the parameters struct
	err := decoder.Decode(&params) // decode into the params
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	// if decode is successful, call the create user function
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	// if creation of user is successful, return user in the response

	respondWithJSON(w, http.StatusCreated, user)

}
