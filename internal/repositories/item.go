package repositories

import (
	"Cart_Api_New/internal/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CartItemRepository struct {
	sql *sqlx.DB
}

func NewCartItemRepository(db *sqlx.DB) *CartItemRepository {
	return &CartItemRepository{sql: db}
}

func (r *CartItemRepository) SaveItem(ctx context.Context, cartItem models.CartItem) (models.CartItem, error) {
	const query = `INSERT INTO items (cart_id, product, quantity ) VALUES ($1, $2, $3)  
	ON CONFLICT (cart_id, product) DO UPDATE SET quantity = items.quantity + EXCLUDED.quantity 
	RETURNING *`

	var dbItem models.CartItem

	err := r.sql.GetContext(ctx, &dbItem, query, cartItem.CartId, cartItem.Product, cartItem.Quantity)
	if err != nil {
		return models.CartItem{}, fmt.Errorf("r.sql.GetContext: %w", err)
	}

	log.Println("new item is successfully created: ", dbItem.Id)
	return dbItem, nil
}

func (r *CartItemRepository) DeleteItem(ctx context.Context, cartItem models.CartItem) error {
	const query = `DELETE FROM items WHERE id = $1 AND cart_id = $2 `
	_, err := r.sql.ExecContext(ctx, query, cartItem.Id, cartItem.CartId)
	if err != nil {
		return fmt.Errorf("r.sql.ExecContext: %w", err)
	}

	log.Println("cart is successfully deleted")
	return nil
}
