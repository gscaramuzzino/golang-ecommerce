package main

import (
	"log"
	"net/http"

	"github.com/gscaramuzzino/go-ecommerce/internal/db"
	"github.com/gscaramuzzino/go-ecommerce/internal/routes"
)

func main() {
	db.Init()
	r := routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
