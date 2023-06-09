package db

import (
	"github.com/gscaramuzzino/go-ecommerce/internal/models"
	"github.com/pkg/errors"
)

func CreateUser(user *models.User) error {
	const sqlQuery = `
			INSERT INTO users (name, email, password)
			VALUES ($1, $2, $3)
			RETURNING id`
	err := db.QueryRow(sqlQuery, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return errors.Wrap(err, "Inserting new user")
	}
	return nil
}
