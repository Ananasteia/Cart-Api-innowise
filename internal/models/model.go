package models

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

func (c *CartItem) Convert() *CartItem {
	return &CartItem{
		Id:       c.Id,
		CartId:   c.CartId,
		Product:  c.Product,
		Quantity: c.Quantity,
	}
}

func (c *Cart) Convert() *Cart {
	items := make([]CartItem, len(c.Items))
	for i, item := range c.Items {
		items[i] = *item.Convert()
	}
	return &Cart{
		Id:    c.Id,
		Items: items,
	}
}
