package main

import (
	"log"
	"net/http"

	"github.com/gscaramuzzino/go-ecommerce/internal/db"
	"github.com/gscaramuzzino/go-ecommerce/internal/models"
	"github.com/gscaramuzzino/go-ecommerce/internal/routes"
)

func main() {
	db.Init()
	seedData()

	r := routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func seedData() {
	products := []models.Product{
		{Name: "Product 1", Description: "This is product 1", Price: 9.99},
		{Name: "Product 2", Description: "This is product 2", Price: 19.99},
		{Name: "Product 3", Description: "This is product 3", Price: 29.99},
	}

	for _, product := range products {
		err := db.CreateProduct(&product)
		if err != nil {
			log.Printf("Failed to seed product: %v. Error: %v", product, err)
		}
	}
}
