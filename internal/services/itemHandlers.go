package services

import (
	"context"
	"fmt"
)

func (a *Service) SaveItem(ctx context.Context, ci CartItem) (*CartItem, error) {
	savedItem, err := a.repo.SaveItem(ctx, ci)
	if err != nil {
		return nil, fmt.Errorf("a.repositories.SaveItem: %w", err)
	}

	return savedItem, nil
}

func (a *Service) DeleteItem(ctx context.Context, ci CartItem) error {
	err := a.repo.DeleteItem(ctx, ci)
	if err != nil {
		return fmt.Errorf("a.repositories.DeleteItem: %w", err)
	}

	return nil
}
