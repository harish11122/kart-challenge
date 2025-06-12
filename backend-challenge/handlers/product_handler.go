package handlers

import (
	"net/http"
	"strconv"

	"backend-challenge/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var products = models.SampleProducts

func ListProducts(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "productId")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid ID supplied"})
		return
	}

	for _, p := range products {
		if p.ID == idStr {
			render.JSON(w, r, p)
			return
		}
	}

	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, map[string]string{"error": "Product not found"})
}
