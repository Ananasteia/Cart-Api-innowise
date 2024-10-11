package handlers

import (
	"Cart_Api_New/internal/services"
	"context"
	"net/http"
	"strings"
)

type appInterface interface {
	SaveItem(ctx context.Context, ci services.CartItem) (*services.CartItem, error)
	DeleteItem(ctx context.Context, ci services.CartItem) error
	GetCart(ctx context.Context, cartId int) (*services.Cart, error)
	CreateNewCart(ctx context.Context) (*services.Cart, error)
}
type api struct {
	app appInterface
}

func New(a appInterface) http.Handler {
	newApi := &api{
		app: a,
	}
	return newApi
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		switch {
		case len(path) == 1 && path[0] == "carts":
			err := a.CreateCart(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		case len(path) == 3 && path[0] == "carts" && path[2] == "items":
			err := a.AddToCart(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	case r.Method == http.MethodGet:
		err := a.ViewCart(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case r.Method == http.MethodDelete:
		err := a.RemoveFromCart(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "rest.ServeHTTP", http.StatusBadRequest)

	}
}
