package repositories

import (
	"Cart_Api_New/internal/models"
	"Cart_Api_New/internal/services"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CartItemRepository struct {
	sql *sqlx.DB
}

func NewCartItemRepository(db *sqlx.DB) *CartItemRepository {
	return &CartItemRepository{sql: db}
}

func (r *CartItemRepository) SaveItem(ctx context.Context, cartItem services.CartItem) (*services.CartItem, error) {
	const query = `insert into items (cart_id, product, quantity ) values ($1, $2, $3) returning *`

	var dbItem models.CartItem

	err := r.sql.GetContext(ctx, &dbItem, query, cartItem.CartId, cartItem.Product, cartItem.Quantity)
	if err != nil {
		return nil, fmt.Errorf("r.sql.GetContext: %w", err)
	}

	return dbItem.Convert(), nil
}

func (r *CartItemRepository) DeleteItem(ctx context.Context, cartItem services.CartItem) error {
	const query = `delete from items WHERE id = $1 and cart_id = $2 `
	_, err := r.sql.ExecContext(ctx, query, cartItem.Id, cartItem.CartId)
	if err != nil {
		return fmt.Errorf("r.sql.ExecContext: %w", err)
	}

	return nil
}
