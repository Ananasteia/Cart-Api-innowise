package repositories

import (
	"Cart_Api_New/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	CartItem
	Cart
}

type (
	CartItem interface {
		SaveItem(ctx context.Context, cartItem models.CartItem) (models.CartItem, error)
		DeleteItem(ctx context.Context, cartItem models.CartItem) error
	}
	Cart interface {
		GetCart(ctx context.Context, cartId int) (models.Cart, error)
		CreateNewCart(ctx context.Context) (models.Cart, error)
	}
)

func New(db *sqlx.DB) Repo {
	return Repo{
		CartItem: NewCartItemRepository(db),
		Cart:     NewCartRepository(db),
	}
}
