package db

import (
	"math/rand"
	"time"

	"github.com/gscaramuzzino/go-ecommerce/internal/models"
)

func CreateProduct(product *models.Product) error {
	const sqlQuery = `
		INSERT INTO products (name, description, price)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(sqlQuery, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func GetAllProducts() ([]*models.Product, error) {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if shouldSimulateDelay() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second * 6) // Adjust the duration as per your requirement
	}

	return products, nil

}

func shouldSimulateDelay() bool {
	// Simulate the condition for when the delay should be introduced

	// Generate a random number between 0 and 9
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	return randomNumber <= 3 || randomNumber >= 97
}

func GetProductByID(id int) (*models.Product, error) {
	row := db.QueryRow("SELECT id, name, description, price FROM products WHERE id = $1", id)

	product := &models.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}

	return product, nil
}
