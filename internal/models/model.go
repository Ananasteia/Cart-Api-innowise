package models

import "Cart_Api_New/internal/services"

type Cart struct {
	Id    int        `db:"id"`
	Items []CartItem `db:"cart_item"`
}

type CartItem struct {
	Id       int    `db:"id"`
	CartId   int    `db:"cart_id"`
	Product  string `db:"product"`
	Quantity int    `db:"quantity"`
}

func (c *CartItem) Convert() *services.CartItem {
	return &services.CartItem{
		Id:       c.Id,
		CartId:   c.CartId,
		Product:  c.Product,
		Quantity: c.Quantity,
	}
}

func (c *Cart) Convert() *services.Cart {
	items := make([]services.CartItem, len(c.Items))
	for i, item := range c.Items {
		items[i] = *item.Convert()
	}
	return &services.Cart{
		Id:    c.Id,
		Items: items,
	}
}
