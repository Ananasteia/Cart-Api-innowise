package handlers

import (
	"Cart_Api_New/internal/errorsx"
	"Cart_Api_New/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (a *api) AddToCart(w http.ResponseWriter, r *http.Request) {
	idNumber, err := strconv.Atoi(r.PathValue("cartId"))
	if err != nil {
		log.Println("from strconv.Atoi: ", err)
		errorWrite(w, errorsx.NumberInvalid, http.StatusNotFound)
		return
	}

	var newItem models.CartItem

	err = json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		log.Println("from json.NewDecoder: ", err)
		errorWrite(w, errorsx.NumberInvalid, http.StatusInternalServerError)
		return
	}

	if newItem.Quantity <= 0 {
		errorWrite(w, errorsx.QuantityInvalid, http.StatusBadRequest)
		return
	}

	newItem.CartId = idNumber

	savedItem, err := a.app.SaveItem(r.Context(), newItem)
	if err != nil {
		log.Println("from a.app.SaveItem: ", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(savedItem)
	if err != nil {
		log.Println("from json.NewEncoder: %w", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	return
}

func (a *api) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	idCartNumber, err := strconv.Atoi(r.PathValue("cartId"))
	if err != nil {
		log.Println("strconv.Atoi: %w", err)
		errorWrite(w, errorsx.NumberInvalid, http.StatusNotFound)
		return
	}

	idItemNumber, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println("strconv.Atoi: %w", err)
		errorWrite(w, errorsx.NumberInvalid, http.StatusNotFound)
		return
	}

	var newItemToDelete models.CartItem

	newItemToDelete.CartId = idCartNumber
	newItemToDelete.Id = idItemNumber

	err = a.app.DeleteItem(r.Context(), newItemToDelete)
	if err != nil {
		log.Println("a.app.DeleteItem: %w", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	return
}
