package jpostcode

import (
	"sync"
)

var (
	adapter     Adapter
	adapterOnce sync.Once
)

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
	adapterOnce.Do(func() {
		// set default adapter
		var err error
		adapter, err = newMapAdapter()
		if err != nil {
			panic(err)
		}
	})
	return adapter.SearchAddressesFromPostCode(postCode)
}
