package services

import (
	"Cart_Api_New/internal/models"
	"context"
	"fmt"
	"log"
)

func (a *Service) CreateNewCart(ctx context.Context) (models.Cart, error) {
	newCart, err := a.repo.CreateNewCart(ctx)
	if err != nil {
		log.Print("a.repo.CreateNewCart: ", err)
		return models.Cart{}, fmt.Errorf("a.repositories.CreateNewCart: %w", err)
	}

	return newCart, nil
}

func (a *Service) GetCart(ctx context.Context, c int) (models.Cart, error) {
	viewedCart, err := a.repo.GetCart(ctx, c)
	if err != nil {
		log.Println("a.repo.GetCart: ", err)
		return models.Cart{}, fmt.Errorf("a.repositories.GetItem: %w", err)
	}

	return viewedCart, nil
}
