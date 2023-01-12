package jpostcode

type Adapter interface {
	SearchAddressesFromPostCode(postCode string) ([]*Address, error)
}
