package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/harish11122/kart-challenge/advanced-challenge/backend-challenge/models"
	"github.com/harish11122/kart-challenge/advanced-challenge/backend-challenge/utils"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateOrderID() string {
	return uuid.NewString()
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid input"})
		return
	}

	if len(req.Items) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Items are required"})
		return
	}

	var orderProducts []models.Product
	for _, item := range req.Items {
		product, ok := getProductByID(item.ProductID)
		if !ok {
			log.Printf("Product %s not found", item.ProductID)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "Invalid product ID"})
			return
		}
		orderProducts = append(orderProducts, product)
	}

	// Validate promo code if present
	var appliedCoupon string
	if req.CouponCode != "" {
		if utils.IsValidPromoCode(req.CouponCode) {
			appliedCoupon = req.CouponCode
		} else {
			render.Status(r, http.StatusUnprocessableEntity)
			render.JSON(w, r, map[string]string{"error": "Invalid promo code"})
			return
		}
	}

	response := models.OrderResponse{
		ID:         generateOrderID(),
		Items:      req.Items,
		Products:   orderProducts,
		CouponCode: appliedCoupon, // Only set if valid
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func getProductByID(id string) (models.Product, bool) {
	products := models.SampleProducts

	for _, p := range products {
		if p.ID == id {
			return p, true
		}
	}
	return models.Product{}, false
}
