package handlers

import (
	"Cart_Api_New/internal/errorsx"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func (h Handler) CreateCart(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), requestTimeToProcess)
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

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeToProcess)
	defer cancel()

	showedCart, err := h.service.GetCart(ctx, idCartNumber)
	switch {
	case errors.Is(err, errorsx.InvalidCartIdErr):
		errorWrite(w, errorsx.InvalidCartIdErr, http.StatusNotFound)
		return
	case errors.Is(err, errorsx.CartNotExistErr):
		errorWrite(w, errorsx.CartNotExistErr, http.StatusNotFound)
		return
	case err != nil:
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
