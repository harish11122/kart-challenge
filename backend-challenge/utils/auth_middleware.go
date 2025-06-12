package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

// testing for sample validation
var ValidAPIKey = "create_order"

func RequireAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api_key")

		if apiKey != ValidAPIKey {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error":   "unauthorized",
				"message": "Invalid or missing API key",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
