package services

import (
	"Cart_Api_New/internal/models"
	"Cart_Api_New/internal/repositories"
	"context"
)

type Servicer interface {
	SaveItem(ctx context.Context, ci models.CartItem) (models.CartItem, error)
	DeleteItem(ctx context.Context, ci models.CartItem) error
	GetCart(ctx context.Context, cartId int) (models.Cart, error)
	CreateNewCart(ctx context.Context) (models.Cart, error)
}

type Service struct {
	repo repositories.Repo
}

func New(r repositories.Repo) *Service {
	return &Service{
		repo: r,
	}
}
