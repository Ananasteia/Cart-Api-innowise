package repositories

import (
	"Cart_Api_New/internal/errorsx"
	"Cart_Api_New/internal/models"
	"context"
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
	var dbItem models.CartItem
	const query = `INSERT INTO items (cart_id, product, quantity ) VALUES ($1, $2, $3)  
	ON CONFLICT (cart_id, product) DO UPDATE SET quantity = items.quantity + EXCLUDED.quantity 
	RETURNING *`
	err := r.sql.GetContext(ctx, &dbItem, query, cartItem.CartId, cartItem.Product, cartItem.Quantity)
	if err != nil {
		return models.CartItem{}, err
	}

	log.Println("new item is successfully created: ", dbItem.Id)
	return dbItem, nil
}

func (r *CartItemRepository) DeleteItem(ctx context.Context, cartItem models.CartItem) error {
	const query = `DELETE FROM items WHERE id = $1 AND cart_id = $2 `
	result, err := r.sql.ExecContext(ctx, query, cartItem.Id, cartItem.CartId)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errorsx.ItemNotExistErr
	}

	log.Println("cart is successfully deleted")
	return nil
}
