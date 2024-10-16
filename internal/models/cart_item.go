package models

type CartItem struct {
	Id       int    `db:"id" json:"id"`
	CartId   int    `db:"cart_id" json:"cartId"`
	Product  string `db:"product" json:"product"`
	Quantity int    `db:"quantity" json:"quantity"`
}
