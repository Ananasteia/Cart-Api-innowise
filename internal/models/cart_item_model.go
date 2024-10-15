package models

type CartItem struct {
	Id       int    `db:"id"`
	CartId   int    `db:"cart_id"`
	Product  string `db:"product"`
	Quantity int    `db:"quantity"`
}
