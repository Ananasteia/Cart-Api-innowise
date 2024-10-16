package models

type Cart struct {
	Id    int        `db:"id" json:"id"`
	Items []CartItem `db:"cart_item" json:"items"`
}
