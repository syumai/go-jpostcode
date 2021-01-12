package jpostcode

func Find(postCode string) (*Address, error) {
	addresses, err := Search(postCode)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, ErrNotFound
	}
	return addresses[0], nil
}

func Search(postCode string) ([]*Address, error) {
	return adapter.SearchAddressesFromPostCode(postCode)
}
