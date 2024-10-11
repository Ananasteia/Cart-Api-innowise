package repositories

import (
	"Cart_Api_New/internal/models"
	"Cart_Api_New/internal/services"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CartRepository struct {
	sql *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{sql: db}
}

func (r *CartRepository) CreateNewCart(ctx context.Context) (*services.Cart, error) {
	const query = `insert into carts default values returning *`

	var dbCart models.Cart

	err := r.sql.GetContext(ctx, &dbCart, query)
	if err != nil {
		return nil, fmt.Errorf("r.sql.GetContext: %w", err)
	}

	return dbCart.Convert(), nil
}

func (r *CartRepository) GetCart(ctx context.Context, cartId int) (*services.Cart, error) {
	const query = `select * from items WHERE cart_id = $1`

	const queryCart = `select * from carts WHERE id = $1`

	var dbCart models.Cart

	err := r.sql.GetContext(ctx, &dbCart, queryCart, cartId)
	if err != nil {
		//	sql.ErrNoRows
		return nil, fmt.Errorf("r.sql.GetContext: %w", err)
	}

	err = r.sql.SelectContext(ctx, &dbCart.Items, query, cartId)
	if err != nil {
		return nil, err
	}

	return dbCart.Convert(), nil
}
