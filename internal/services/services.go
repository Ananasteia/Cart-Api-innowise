package services

import (
	"Cart_Api_New/internal/models"
	"context"
)

type (
	Service struct {
		repo RepoInterface
	}
	RepoInterface interface {
		SaveItem(ctx context.Context, cartItem models.CartItem) (models.CartItem, error)
		DeleteItem(ctx context.Context, cartItem models.CartItem) error
		GetCart(ctx context.Context, cartId int) (models.Cart, error)
		CreateNewCart(ctx context.Context) (models.Cart, error)
	}
)

func New(r RepoInterface) *Service {
	return &Service{
		repo: r,
	}
}
