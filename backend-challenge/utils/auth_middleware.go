package utils

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// testing for sample validation
var ValidAPIKey = "apitest"

func RequireAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api_key")
		log.Printf("test", r.Header)

		if apiKey != ValidAPIKey {
			log.Printf("Unauthorized request", "provided_key", apiKey)
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
