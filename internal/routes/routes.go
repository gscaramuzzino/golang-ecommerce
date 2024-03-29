package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gscaramuzzino/go-ecommerce/internal/handlers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	// Product routes
	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")
	r.HandleFunc("/products/{productID}", handlers.GetProductDetails).Methods("GET")

	// Cart routes
	r.HandleFunc("/cart", handlers.ViewCart).Methods("GET")
	r.HandleFunc("/cart/add/{productID}", handlers.AddToCart).Methods("POST")

	// Product routes
	r.HandleFunc("/cart/checkout", handlers.Checkout).Methods("POST")

	r.HandleFunc("/swagger-ui/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./dist"))
	}).Methods("GET")

	return r
}
