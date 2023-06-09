package db

import "github.com/gscaramuzzino/go-ecommerce/internal/models"

func AddProductToCart(userID int, productID int) error {
	_, err := db.Exec(`
		INSERT INTO carts (user_id, product_id) 
		VALUES ($1, $2)`,
		userID,
		productID,
	)
	return err
}

func GetUserCart(userID int) (*models.Cart, error) {
	rows, err := db.Query(`
		SELECT p.id, p.name, p.description, p.price
		FROM products p
		JOIN carts c ON p.id = c.product_id
		WHERE c.user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cart := &models.Cart{UserID: userID}
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		cart.Products = append(cart.Products, product)
	}

	return cart, nil
}
