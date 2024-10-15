package repositories

import (
	"Cart_Api_New/internal/errorsx"
	"Cart_Api_New/internal/models"
	"context"
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
		return models.Cart{}, err
	}

	log.Println("new cart is successfully created: ", dbCart.Id)
	return dbCart, nil
}

func (r *CartRepository) GetCart(ctx context.Context, cartId int) (models.Cart, error) {
	err := r.checkExistence(ctx, cartId)
	if err != nil {
		return models.Cart{}, err
	}

	const query = `SELECT carts.id, items.id, product, quantity FROM carts LEFT JOIN items 
    ON items.cart_id = carts.id WHERE carts.id = $1`

	var dbCart models.Cart

	err = r.sql.SelectContext(ctx, &dbCart.Items, query, cartId)
	if err != nil {
		return models.Cart{}, err
	}

	dbCart.Id = cartId

	log.Println("cart is successfully got: ", dbCart.Id)
	return dbCart, nil
}

func (r *CartRepository) checkExistence(ctx context.Context, cartId int) error {
	const query = `SELECT id FROM carts WHERE id = $1`
	var checkingValue int
	err := r.sql.GetContext(ctx, &checkingValue, query, cartId)
	if err != nil {
		return errorsx.InvalidCartIdErr
	}
	return nil
}
