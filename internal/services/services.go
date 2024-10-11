package services

import (
	"context"
)

type (
	Cart struct {
		Id    int
		Items []CartItem
	}
	CartItem struct {
		Id       int    `json:"id"`
		CartId   int    `json:"cart_id"`
		Product  string `json:"product"`
		Quantity int    `json:"quantity"`
	}
	Service struct {
		repo RepoInterface
	}
	RepoInterface interface {
		SaveItem(ctx context.Context, cartItem CartItem) (*CartItem, error)
		DeleteItem(ctx context.Context, cartItem CartItem) error
		GetCart(ctx context.Context, cartId int) (*Cart, error)
		CreateNewCart(ctx context.Context) (*Cart, error)
	}
)

func New(r RepoInterface) *Service {
	return &Service{
		repo: r,
	}
}
