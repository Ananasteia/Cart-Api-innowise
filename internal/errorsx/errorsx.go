package errorsx

import (
	"errors"
)

var (
	StatusServerErr = errors.New("OOPS, something went wrong")
	NumberInvalid   = errors.New("not valid info, try again")
	QuantityInvalid = errors.New("quantity couldn't be non positive")
)
