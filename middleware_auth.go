package main

import (
	"fmt"
	"net/http"

	"github.com/kghsachin/learn_go/internal/auth"
	"github.com/kghsachin/learn_go/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err!=nil{
			respondWithError(w, 400, fmt.Sprintf("Auth error : %s", err))
			return 
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err!=nil {
			respondWithError(w, 400, fmt.Sprintf("Error getting user: %s", err))
			return 
		}
		handler(w, r, user)
	}
}