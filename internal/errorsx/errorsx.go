package errorsx

import (
	"errors"
)

var (
	InternalServerErr  = errors.New("OOPS, something went wrong")
	InvalidIdErr       = errors.New("not valid item id, try again")
	InvalidCartIdErr   = errors.New("not valid cart id, try again")
	InvalidQuantityErr = errors.New("quantity should be positive")
	ItemNotExistErr    = errors.New("item doesn't exist in this cart")
	CartNotExistErr    = errors.New("cart doesn't exist")
)
