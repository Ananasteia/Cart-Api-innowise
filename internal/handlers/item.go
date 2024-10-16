package handlers

import (
	"Cart_Api_New/internal/errorsx"
	"Cart_Api_New/internal/models"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func (h Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	idNumber, err := strconv.Atoi(r.PathValue("cartId"))
	if err != nil {
		log.Printf("from strconv.Atoi: %v", err)
		errorWrite(w, errorsx.InvalidCartIdErr, http.StatusBadRequest)
		return
	}

	var newItem models.CartItem

	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		log.Printf("from json.NewDecoder: %v", err)
		errorWrite(w, errorsx.InvalidCartIdErr, http.StatusInternalServerError)
		return
	}

	if newItem.Quantity <= 0 {
		errorWrite(w, errorsx.InvalidQuantityErr, http.StatusBadRequest)
		return
	}

	newItem.CartId = idNumber

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeToProcess)
	defer cancel()

	savedItem, err := h.service.SaveItem(ctx, newItem)
	if err != nil {
		log.Printf("from h.service.SaveItem: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(savedItem)
	if err != nil {
		log.Printf("from json.NewEncoder: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}
}

func (h Handler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	idCartNumber, err := strconv.Atoi(r.PathValue("cartId"))
	if err != nil {
		log.Printf("strconv.Atoi: %v", err)
		errorWrite(w, errorsx.InvalidCartIdErr, http.StatusBadRequest)
		return
	}

	idItemNumber, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("strconv.Atoi: %v", err)
		errorWrite(w, errorsx.InvalidIdErr, http.StatusBadRequest)
		return
	}

	var newItemToDelete models.CartItem

	newItemToDelete.CartId = idCartNumber
	newItemToDelete.Id = idItemNumber

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeToProcess)
	defer cancel()

	err = h.service.DeleteItem(ctx, newItemToDelete)
	switch {
	case errors.Is(err, errorsx.ItemNotExistErr):
		errorWrite(w, errorsx.ItemNotExistErr, http.StatusNotFound)
		return
	case err != nil:
		log.Printf("h.service.DeleteItem: %v", err)
		errorWrite(w, errorsx.InternalServerErr, http.StatusInternalServerError)
		return
	}
}
