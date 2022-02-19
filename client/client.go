package client

import (
	"net/http"
	"time"
)

type Product struct {
	ID          int       `json: "id"`
	Name        string    `json: "name"`
	Description string    `json: "description"`
	Price       float32   `json: "price"`
	SKU         string    `json: "sku"`
	CreatedAt   time.Time `json: "-"`
	UpdatedAt   time.Time `json: "-"`
}

type Products []*Product

var httpClient = &http.Client{}
