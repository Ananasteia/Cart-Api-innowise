package repositories

import (
	"Cart_Api_New/internal/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CartRepository struct {
	sql *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{sql: db}
}

func (r *CartRepository) CreateNewCart(ctx context.Context) (models.Cart, error) {
	const query = `INSERT INTO carts DEFAULT VALUES returning *`

	var dbCart models.Cart

	err := r.sql.GetContext(ctx, &dbCart, query)
	if err != nil {
		return models.Cart{}, fmt.Errorf("r.sql.GetContext: %w", err)
	}

	log.Println("new cart is successfully created: ", dbCart.Id)
	return dbCart, nil
}

func (r *CartRepository) GetCart(ctx context.Context, cartId int) (models.Cart, error) {
	const query = `SELECT * FROM items WHERE cart_id = $1`

	const queryCart = `SELECT * FROM carts WHERE id = $1`

	var dbCart models.Cart

	err := r.sql.GetContext(ctx, &dbCart, queryCart, cartId)
	if err != nil {
		return models.Cart{}, fmt.Errorf("r.sql.GetContext: %w", err)
	}

	err = r.sql.SelectContext(ctx, &dbCart.Items, query, cartId)
	if err != nil {
		return models.Cart{}, err
	}

	log.Println("cart is successfully got: ", dbCart.Id)
	return dbCart, nil
}
