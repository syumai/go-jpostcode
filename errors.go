package jpostcode

import "fmt"

var (
	ErrNotFound        = fmt.Errorf("postcode is not found")
	ErrInvalidArgument = fmt.Errorf("postcode is invalid")
	ErrInternal        = fmt.Errorf("internal error")
)
