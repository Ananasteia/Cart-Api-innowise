package handlers

import (
	"Cart_Api_New/internal/services"
	"log"
	"net/http"
)

type Handler struct {
	service services.Servicer
}

func (h Handler) Handle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /carts", h.CreateCart)
	mux.HandleFunc("GET /carts/{id}", h.GetCart)
	mux.HandleFunc("DELETE /carts/{cartId}/items/{id}", h.RemoveFromCart)
	mux.HandleFunc("POST /carts/{cartId}/items", h.AddToCart)
	return mux
}

func New(s services.Servicer) Handler {
	return Handler{
		service: s,
	}
}

func errorWrite(w http.ResponseWriter, text error, statusCode int) {
	http.Error(w, text.Error(), statusCode)

	log.Println(statusCode, text)
}
