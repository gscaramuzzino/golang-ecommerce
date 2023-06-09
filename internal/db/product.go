package db

import "github.com/gscaramuzzino/go-ecommerce/internal/models"

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

	return products, nil
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
