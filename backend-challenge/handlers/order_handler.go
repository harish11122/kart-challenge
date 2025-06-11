package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
	"github.com/harish11122/kart-challenge/advanced-challenge/backend-challenge/models"
	"github.com/harish11122/kart-challenge/advanced-challenge/backend-challenge/utils"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateOrderID() string {
	const chars = "0123456789abcdef"
	res := make([]byte, 16)
	for i := range res {
		res[i] = chars[rand.Intn(len(chars))]
	}
	return strings.Join(chunks(res, 4), "-")
}

func chunks(slice []byte, chunkSize int) []string {
	var chunks [][]byte
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	strs := make([]string, len(chunks))
	for i, c := range chunks {
		strs[i] = string(c)
	}
	return strs
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req models.OrderRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid input"})
		return
	}
	log.Printf("", req)
	log.Printf("", req.Items)
	log.Printf("", req.CouponCode)
	if len(req.Items) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Items are required"})
		return
	}

	if req.CouponCode != "" && !utils.IsValidPromoCode(req.CouponCode) {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, map[string]string{"error": "Invalid promo code"})
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

	response := models.OrderResponse{
		ID:       generateOrderID(),
		Items:    req.Items,
		Products: orderProducts,
	}

	render.JSON(w, r, response)
}

func getProductByID(id string) (models.Product, bool) {
	products := []models.Product{
		{ID: "1", Name: "Chicken Waffle", Price: 9.99, Category: "Waffle"},
		{ID: "2", Name: "Veggie Burger", Price: 7.50, Category: "Burger"},
		{ID: "3", Name: "Fries", Price: 2.99, Category: "Side"},
	}

	for _, p := range products {
		if p.ID == id {
			return p, true
		}
	}
	return models.Product{}, false
}
