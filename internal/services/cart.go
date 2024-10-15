package services

import (
	"Cart_Api_New/internal/models"
	"context"
	"log"
)

func (a *Service) CreateNewCart(ctx context.Context) (models.Cart, error) {
	newCart, err := a.repo.CreateNewCart(ctx)
	if err != nil {
		log.Printf("a.repo.CreateNewCart: %v", err)
		return models.Cart{}, err
	}

	return newCart, nil
}

func (a *Service) GetCart(ctx context.Context, idCart int) (models.Cart, error) {
	viewedCart, err := a.repo.GetCart(ctx, idCart)
	if err != nil {
		log.Printf("a.repo.GetCart: %v", err)
		return models.Cart{}, err
	}

	return viewedCart, nil
}
