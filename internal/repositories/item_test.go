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

func TestCartItemRepository_SaveItem(t *testing.T) {
	testStruct := []struct {
		name         string
		cartItem     models.CartItem
		mockSetup    func(mock sqlmock.Sqlmock, input models.CartItem)
		expectedItem models.CartItem
		expectedErr  error
	}{
		{
			name: "created item",
			cartItem: models.CartItem{
				CartId:   1,
				Product:  "something",
				Quantity: 10,
			},
			mockSetup: func(mock sqlmock.Sqlmock, input models.CartItem) {
				mock.ExpectQuery("INSERT INTO items").
					WithArgs(input.CartId, input.Product, input.Quantity).
					WillReturnRows(sqlmock.NewRows([]string{"cart_id", "product", "quantity"}).AddRow(input.CartId, input.Product, input.Quantity))
			},
			expectedItem: models.CartItem{
				CartId:   1,
				Product:  "something",
				Quantity: 10,
			},
			expectedErr: nil,
		}, {
			name: "updated item",
			cartItem: models.CartItem{
				CartId:   1,
				Product:  "something",
				Quantity: 5,
			},
			mockSetup: func(mock sqlmock.Sqlmock, input models.CartItem) {
				mock.ExpectQuery("INSERT INTO items").
					WithArgs(input.CartId, input.Product, input.Quantity).
					WillReturnRows(sqlmock.NewRows([]string{"cart_id", "product", "quantity"}).AddRow(input.CartId, input.Product, input.Quantity+10))
			},
			expectedItem: models.CartItem{
				CartId:   1,
				Product:  "something",
				Quantity: 15,
			},
			expectedErr: nil,
		},
	}
	for _, testedVariant := range testStruct {
		t.Run(testedVariant.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error creating sqlmock: %s", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "132")
			repo := repositories.New(sqlxDB)

			testedVariant.mockSetup(mock, testedVariant.cartItem)

			resultItem, err := repo.SaveItem(context.Background(), testedVariant.cartItem)

			assert.Equal(t, testedVariant.expectedItem, resultItem)
			assert.Equal(t, testedVariant.expectedErr, err)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestCartItemRepository_DeleteItem(t *testing.T) {
	testStruct := []struct {
		name        string
		cartItem    models.CartItem
		mockSetup   func(mock sqlmock.Sqlmock)
		expectedErr error
	}{
		{name: "deleted",
			cartItem: models.CartItem{
				Id:       1,
				CartId:   1,
				Product:  "something",
				Quantity: 15,
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM items").
					WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
	}
	for _, testedVariant := range testStruct {
		t.Run(testedVariant.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error creating sqlmock: %s", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "132")
			repo := repositories.New(sqlxDB)

			testedVariant.mockSetup(mock)

			err = repo.DeleteItem(context.Background(), testedVariant.cartItem)

			assert.Equal(t, err, testedVariant.expectedErr)

			err = mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
