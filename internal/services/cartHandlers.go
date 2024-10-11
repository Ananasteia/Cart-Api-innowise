package services

import (
	"context"
	"fmt"
	"log"
)

func (a *Service) CreateNewCart(ctx context.Context) (*Cart, error) {
	newCart, err := a.repo.CreateNewCart(ctx)
	if err != nil {
		log.Print("CreateNewCart")
		return nil, fmt.Errorf("a.repositories.CreateNewCart: %w", err)
	}

	return newCart, nil
}

func (a *Service) GetCart(ctx context.Context, c int) (*Cart, error) {
	viewedCart, err := a.repo.GetCart(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("a.repositories.GetItem: %w", err)
	}

	return viewedCart, nil
}
