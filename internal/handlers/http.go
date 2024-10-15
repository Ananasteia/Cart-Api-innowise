package handlers

import (
	"Cart_Api_New/internal/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type appInterface interface {
	SaveItem(ctx context.Context, ci models.CartItem) (models.CartItem, error)
	DeleteItem(ctx context.Context, ci models.CartItem) error
	GetCart(ctx context.Context, cartId int) (models.Cart, error)
	CreateNewCart(ctx context.Context) (models.Cart, error)
}
type api struct {
	app appInterface
}

func New(a appInterface) *http.ServeMux {
	newApi := &api{
		app: a,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /carts", newApi.CreateCart)
	mux.HandleFunc("GET /carts/{id}", newApi.GetCart)
	mux.HandleFunc("DELETE /carts/{cartId}/items/{id}", newApi.RemoveFromCart)
	mux.HandleFunc("POST /carts/{cartId}/items", newApi.AddToCart)
	return mux

}

func errorWrite(w http.ResponseWriter, text error, statusCode int) {

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(text)
	if err != nil {
		log.Println(err)
	}
	log.Println(statusCode, text)
}
