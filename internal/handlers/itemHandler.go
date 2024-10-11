package handlers

import (
	"Cart_Api_New/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (a *api) AddToCart(w http.ResponseWriter, r *http.Request) error {
	path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	idNumber, err := strconv.Atoi(path[1])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	var newItem models.CartItem

	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		return fmt.Errorf("json.NewDecoder: %w", err)
	}

	if newItem.Quantity <= 0 {
		return fmt.Errorf("Quantity couldn't be non positive ")
	}

	newItem.CartId = idNumber

	savedItem, err := a.app.SaveItem(r.Context(), newItem)
	if err != nil {
		return fmt.Errorf("a.app.SaveItem: %w", err)
	}

	err = json.NewEncoder(w).Encode(savedItem)
	if err != nil {
		return fmt.Errorf("json.NewEncoder: %w", err)
	}

	return nil
}

func (a *api) RemoveFromCart(w http.ResponseWriter, r *http.Request) error {
	path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	idCartNumber, err := strconv.Atoi(path[1])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	idItemNumber, err := strconv.Atoi(path[3])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	var newItemToDelete models.CartItem

	newItemToDelete.CartId = idCartNumber
	newItemToDelete.Id = idItemNumber

	err = a.app.DeleteItem(r.Context(), newItemToDelete)
	if err != nil {
		return fmt.Errorf("a.app.DeleteItem: %w", err)
	}

	return nil
}
