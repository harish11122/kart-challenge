package main

import (
	"log"
	"net/http"

	"backend-challenge/handlers"
	"backend-challenge/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Product routes
	r.Group(func(r chi.Router) {
		r.Get("/product", handlers.ListProducts)
		r.Get("/product/{productId}", handlers.GetProduct)
	})

	// Order routes
	r.Group(func(r chi.Router) {
		r.Use(utils.RequireAPIKey)
		r.Post("/order", handlers.PlaceOrder)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
