package repositories

import (
	"Cart_Api_New/internal/services"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	CartItem
	Cart
}

type (
	CartItem interface {
		SaveItem(ctx context.Context, cartItem services.CartItem) (*services.CartItem, error)
		DeleteItem(ctx context.Context, cartItem services.CartItem) error
	}
	Cart interface {
		GetCart(ctx context.Context, cartId int) (*services.Cart, error)
		CreateNewCart(ctx context.Context) (*services.Cart, error)
	}
)

func New(db *sqlx.DB) Repo {
	return Repo{
		CartItem: NewCartItemRepository(db),
		Cart:     NewCartRepository(db),
	}
}
