package services

import (
	"Cart_Api_New/internal/models"
	"context"
	"log"
)

func (a *Service) SaveItem(ctx context.Context, cartItem models.CartItem) (models.CartItem, error) {
	savedItem, err := a.repo.SaveItem(ctx, cartItem)
	if err != nil {
		log.Printf("a.repo.DeleteItem: %v", err)
		return models.CartItem{}, err
	}

	return savedItem, nil
}

func (a *Service) DeleteItem(ctx context.Context, cartItem models.CartItem) error {
	err := a.repo.DeleteItem(ctx, cartItem)
	if err != nil {
		log.Printf("a.repo.DeleteItem: %v", err)
		return err
	}

	return nil
}
