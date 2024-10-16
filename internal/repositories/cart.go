package repositories

import (
	"Cart_Api_New/internal/errorsx"
	"Cart_Api_New/internal/models"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type CartRepository struct {
	sql *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{sql: db}
}

func (r *CartRepository) CreateNewCart(ctx context.Context) (models.Cart, error) {
	var dbCart models.Cart
	const query = `INSERT INTO carts DEFAULT VALUES returning *`
	err := r.sql.GetContext(ctx, &dbCart, query)
	if err != nil {
		return models.Cart{}, err
	}

	log.Println("new cart is successfully created: ", dbCart.Id)
	return dbCart, nil
}

func (r *CartRepository) GetCart(ctx context.Context, cartId int) (models.Cart, error) {
	err := r.checkExistence(ctx, cartId)
	if err != nil {
		return models.Cart{}, err
	}

	var testDBIem []dbItem
	const query = `SELECT * FROM items WHERE cart_id = $1`
	err = r.sql.SelectContext(ctx, &testDBIem, query, cartId)
	if err != nil {
		return models.Cart{}, err
	}

	dbCart := models.Cart{
		Items: make([]models.CartItem, 0, len(testDBIem)),
	}
	for _, item := range testDBIem {
		dbCart.Items = append(dbCart.Items, item.convert())
	}

	dbCart.Id = cartId

	log.Println("cart is successfully got: ", dbCart.Id)
	return dbCart, nil
}

type dbItem struct {
	Id       sql.NullInt32  `db:"id"`
	CartId   sql.NullInt32  `db:"cart_id"`
	Product  sql.NullString `db:"product"`
	Quantity sql.NullInt32  `db:"quantity"`
}

func (dbI dbItem) convert() models.CartItem {
	var res models.CartItem
	if dbI.Id.Valid {
		res.Id = int(dbI.Id.Int32)
	}
	if dbI.CartId.Valid {
		res.CartId = int(dbI.CartId.Int32)
	}
	if dbI.Product.Valid {
		res.Product = dbI.Product.String
	}
	if dbI.Quantity.Valid {
		res.Quantity = int(dbI.Quantity.Int32)
	}
	return res

}

func (r *CartRepository) checkExistence(ctx context.Context, cartId int) error {
	var checkingValue int
	const query = `SELECT id FROM carts WHERE id = $1`
	err := r.sql.GetContext(ctx, &checkingValue, query, cartId)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errorsx.CartNotExistErr
	case err != nil:
		return err
	}

	return nil
}
