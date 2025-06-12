package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-challenge/handlers"
	"backend-challenge/models"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestPlaceOrder_ValidCoupon(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/order", handlers.PlaceOrder)

	server := httptest.NewServer(r)
	defer server.Close()

	orderReq := models.OrderRequest{
		Items: []models.OrderItem{
			{ProductID: "1", Quantity: 2},
		},
		CouponCode: "HAPPYHRS",
	}

	body, _ := json.Marshal(orderReq)

	req, _ := http.NewRequest("POST", server.URL+"/order", bytes.NewBuffer(body))
	req.Header.Set("api_key", "apitest")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPlaceOrder_InvalidCoupon(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/order", handlers.PlaceOrder)

	server := httptest.NewServer(r)
	defer server.Close()

	orderReq := models.OrderRequest{
		Items: []models.OrderItem{
			{ProductID: "1", Quantity: 2},
		},
		CouponCode: "INVALIDCODE",
	}

	body, _ := json.Marshal(orderReq)

	req, _ := http.NewRequest("POST", server.URL+"/order", bytes.NewBuffer(body))
	req.Header.Set("api_key", "create_order")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}
