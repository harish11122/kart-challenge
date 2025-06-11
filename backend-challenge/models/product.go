package models

type ProductImage struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

type Product struct {
	ID       string       `json:"id"`
	Image    ProductImage `json:"image"`
	Name     string       `json:"name"`
	Category string       `json:"category"`
	Price    float64      `json:"price"`
}
