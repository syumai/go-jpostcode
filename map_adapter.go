package jpostcode

import (
	"compress/gzip"
	"encoding/gob"
)

type mapAdapter map[string][]*Address

func newMapAdapter() (mapAdapter, error) {
	f, err := staticFS.Open("data/map.gob.gz")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	var m mapAdapter
	gdec := gob.NewDecoder(gr)
	if err := gdec.Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}

func (a mapAdapter) SearchAddressesFromPostCode(postCode string) ([]*Address, error) {
	if len(postCode) != 7 {
		return nil, ErrInvalidArgument
	}

	addresses, ok := a[postCode]
	if !ok {
		return nil, ErrNotFound
	}
	return addresses, nil
}
