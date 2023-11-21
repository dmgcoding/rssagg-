package main

import (
	"fmt"
	"net/http"

	"github.com/dmgcoding/rssagg/internal/auth"
	"github.com/dmgcoding/rssagg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), api_key)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
			return
		}
		handler(w, r, user)
	}
}
