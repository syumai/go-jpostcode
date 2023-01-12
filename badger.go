package jpostcode

import (
	"bytes"
	"compress/gzip"
	"embed"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/dgraph-io/badger/v3"
	"github.com/mitchellh/mapstructure"
)

//go:embed badger/*
var staticFS embed.FS

type badgerAdapter struct {
	db *badger.DB
}

func newBadgerAdapter() (*badgerAdapter, error) {
	f, err := staticFS.Open("badger/dump.db.gz")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	opt := badger.DefaultOptions("").WithInMemory(true)
	opt.Logger = nil

	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	err = db.Load(gr, 100)
	if err != nil {
		return nil, err
	}

	return &badgerAdapter{
		db: db,
	}, nil
}

func (a *badgerAdapter) SearchAddressesFromPostCode(postCode string) ([]*Address, error) {
	if len(postCode) != 7 {
		return nil, ErrInvalidArgument
	}

	if a.db.IsClosed() {
		return nil, fmt.Errorf("badger db is already closed: %w", ErrInternal)
	}

	var addressData interface{}
	err := a.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(postCode))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			buf := bytes.NewBuffer(val)
			return json.NewDecoder(buf).Decode(&addressData)
		})
	})

	if err != nil {
		switch err {
		case badger.ErrKeyNotFound:
			return nil, ErrNotFound
		}
		return nil, err
	}

	var addresses []*Address
	switch reflect.TypeOf(addressData).Kind() {
	case reflect.Slice:
		rawAddrs, ok := addressData.([]interface{})
		if !ok {
			return nil, ErrInternal
		}
		for _, rawAddr := range rawAddrs {
			addr, err := convertJSONToAddress(rawAddr)
			if err != nil {
				return nil, err
			}
			addresses = append(addresses, addr)
		}
	default:
		addr, err := convertJSONToAddress(addressData)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}
	return addresses, nil
}

func convertJSONToAddress(input interface{}) (*Address, error) {
	var addr Address
	err := mapstructure.Decode(input, &addr)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}
