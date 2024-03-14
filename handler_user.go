package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	
	"github.com/kghsachin/learn_go/internal/database"
)

// function signature to define http in go
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing data : %s", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %s", err))
		return
	}
	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
  respondWithJson(w, 200, databaseUserToUser(user))
	// apiKey, err := auth.GetAPIKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("Auth error : %s", err))
	// 	return
	// }
	// user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

	// if err != nil {
	// 	respondWithError(w, 400, fmt.Sprintf("Error getting user: %s", err))
	// 	return
	// }
	// respondWithJson(w, 200, databaseUserToUser(user))

}
