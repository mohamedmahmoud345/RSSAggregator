package main

import (
	"fmt"
	"net/http"

	"github.com/mohamedmahmoud345/RssAggregator/internal/auth"
	"github.com/mohamedmahmoud345/RssAggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func(apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("auth error: %v", err))
		return
	}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("couldn't get user: %v", err))
			return 
		}
		handler(w, r, user)
	}
}

