package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/sfx09/woodchuck/internal/database"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, database.User)

func (c *Controller) HandleAuthentication(next AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetAPIKey(r.Header)
		if err != nil {
			respondWithErr(w, http.StatusUnauthorized, err.Error())
		}
		user, err := c.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithErr(w, http.StatusNotFound, "User not found!")
		}
		next(w, r, user)
	}
}

func GetAPIKey(h http.Header) (string, error) {
	authHeader := h.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("No authorization header included")
	}
	authTokens := strings.Split(authHeader, " ")
	if len(authTokens) < 2 || authTokens[0] != "ApiKey" {
		return "", errors.New("Malformed authorization header")
	}
	return authTokens[1], nil
}
