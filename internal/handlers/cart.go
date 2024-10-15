package handlers

import (
	"Cart_Api_New/internal/errorsx"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) CreateCart(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*1)

	defer cancel()

	newCart, err := h.service.CreateNewCart(ctx)

	if err != nil {
		log.Printf("from h.service.CreateNewCart: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(newCart)
	if err != nil {
		log.Printf("from json.NewEncoder: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}
}

func (h Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	idCartNumber, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("from strconv.Atoi: %v", err)
		errorWrite(w, errorsx.InvalidCartIdErr, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*1)

	defer cancel()

	showedCart, err := h.service.GetCart(ctx, idCartNumber)
	if errors.Is(err, errorsx.InvalidCartIdErr) {
		errorWrite(w, errorsx.InvalidCartIdErr, http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("from h.service.GetCart: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(showedCart)
	if err != nil {
		log.Printf("json.NewEncoder: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}

	return
}
