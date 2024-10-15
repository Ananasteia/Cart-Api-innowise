package repositories_test

import (
	"Cart_Api_New/internal/models"
	"Cart_Api_New/internal/repositories"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCartRepository_CreateNewCart(t *testing.T) {
	testStruct := []struct {
		query        string
		name         string
		cartId       int
		mockSetup    func(mock sqlmock.Sqlmock, query string, cartId int)
		expectedCart models.Cart
		expectedErr  error
	}{
		{
			query:  `INSERT INTO carts DEFAULT VALUES returning *`,
			name:   "create first cart",
			cartId: 1,
			mockSetup: func(mock sqlmock.Sqlmock, query string, cartId int) {
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(cartId))
			},
			expectedCart: models.Cart{Id: 1},
			expectedErr:  nil,
		}, {
			query:  `INSERT INTO carts DEFAULT VALUES returning *`,
			name:   "create second cart",
			cartId: 2,
			mockSetup: func(mock sqlmock.Sqlmock, query string, cartId int) {
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(cartId))
			},
			expectedCart: models.Cart{Id: 2},
			expectedErr:  nil,
		}, {
			query:  `INSERT INTO carts DEFAULT VALUES returning *`,
			name:   "create third cart",
			cartId: 3,
			mockSetup: func(mock sqlmock.Sqlmock, query string, cartId int) {
				mock.ExpectQuery(query).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(cartId))
			},
			expectedCart: models.Cart{Id: 3},
			expectedErr:  nil,
		},
	}
	for _, testedVariant := range testStruct {
		t.Run(testedVariant.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			//defer db.Close()

			sqlxDB := sqlx.NewDb(db, "123")
			repo := repositories.New(sqlxDB)

			testedVariant.mockSetup(mock, testedVariant.query, testedVariant.cartId)

			res, err := repo.Cart.CreateNewCart(context.Background())

			assert.Equal(t, testedVariant.expectedCart, res)
			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
			assert.Equal(t, testedVariant.expectedErr, err)

		})
	}
}

func TestCartRepository_GetCart(t *testing.T) {
	testStruct := []struct {
		name         string
		cartId       int
		mockSetup    func(mock sqlmock.Sqlmock, cartId int)
		expectedCart models.Cart
		expectedErr  error
	}{
		{
			name:   "Empty Cart",
			cartId: 1,
			mockSetup: func(mock sqlmock.Sqlmock, cartId int) {

				mock.ExpectQuery(`SELECT \* FROM carts WHERE id = \$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery(`SELECT \* FROM items WHERE cart_id = \$1`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "cart_id", "product", "quantity"}))
			},
			expectedCart: models.Cart{
				Id:    1,
				Items: []models.CartItem(nil)},
			expectedErr: nil,
		},
	}
	for _, testedVariant := range testStruct {
		t.Run(testedVariant.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			sqlxDB := sqlx.NewDb(db, "123")
			repo := repositories.New(sqlxDB)
			testedVariant.mockSetup(mock, testedVariant.cartId)

			resultCart, err := repo.GetCart(context.Background(), testedVariant.cartId)

			assert.NoError(t, err)
			assert.Equal(t, testedVariant.expectedCart, resultCart)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
