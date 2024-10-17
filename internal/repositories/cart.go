package repositories

import (
	"Cart_Api_New/internal/errorsx"
	"Cart_Api_New/internal/models"
	"context"
	"database/sql"
	"errors"
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
	var dbCart models.Cart
	const query = `INSERT INTO carts DEFAULT VALUES returning *`
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

	var testDBItem []models.CartItem
	const query = `SELECT * FROM carts JOIN items ON carts.id=items.cart_id WHERE carts.id = $1`
	err = r.sql.SelectContext(ctx, &testDBItem, query, cartId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.Cart{Id: cartId, Items: []models.CartItem{}}, nil
		default:
			return models.Cart{}, err
		}
	}

	return models.Cart{
		Id:    cartId,
		Items: testDBItem,
	}, nil
}

func (r *CartRepository) checkExistence(ctx context.Context, cartId int) error {
	var checkingValue int
	const query = `SELECT id FROM carts WHERE id = $1`
	err := r.sql.GetContext(ctx, &checkingValue, query, cartId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errorsx.CartNotExistErr
		default:
			return err
		}
	}
	return nil
}
