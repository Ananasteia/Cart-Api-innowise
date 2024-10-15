package handlers

import (
	"Cart_Api_New/internal/errorsx"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (a *api) CreateCart(w http.ResponseWriter, r *http.Request) {
	newCart, err := a.app.CreateNewCart(r.Context())
	if err != nil {
		log.Println("from a.app.CreateNewCart: ", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(newCart)
	if err != nil {
		log.Println("from json.NewEncoder: %w", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	return
}

func (a *api) GetCart(w http.ResponseWriter, r *http.Request) {
	idCartNumber, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println("from strconv.Atoi: %w", err)
		errorWrite(w, errorsx.NumberInvalid, http.StatusNotFound)
		return
	}

	showedCart, err := a.app.GetCart(r.Context(), idCartNumber)
	if err != nil {
		log.Println("from a.app.GetCart: %w", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(showedCart)
	if err != nil {
		log.Println("json.NewEncoder: %w", err)
		errorWrite(w, errorsx.StatusServerErr, http.StatusInternalServerError)
		return
	}

	return
}
