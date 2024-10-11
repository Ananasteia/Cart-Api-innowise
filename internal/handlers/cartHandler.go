package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (a *api) CreateCart(w http.ResponseWriter, r *http.Request) error {
	newCart, err := a.app.CreateNewCart(r.Context())
	if err != nil {
		return fmt.Errorf("a.app.CreateNewCart: %w", err)
	}

	err = json.NewEncoder(w).Encode(newCart)
	if err != nil {
		return fmt.Errorf("json.NewEncoder: %w", err)
	}

	return nil
}

func (a *api) GetCart(w http.ResponseWriter, r *http.Request) error {
	path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	idCartNumber, err := strconv.Atoi(path[1])
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}

	showedCart, err := a.app.GetCart(r.Context(), idCartNumber)
	if err != nil {
		return fmt.Errorf("a.app.GetCart: %w", err)
	}

	err = json.NewEncoder(w).Encode(showedCart)
	if err != nil {
		return fmt.Errorf("json.NewEncoder: %w", err)
	}

	return nil
}
