package jpostcode

var adapter Adapter

func init() {
	// set default adapter
	var err error
	adapter, err = newBadgerAdapter()
	if err != nil {
		panic(err)
	}
	// closing DB is not needed because badger adapter is using in-memory DB
	// adapter.db.Close()
}

type Adapter interface {
	SearchAddressesFromPostCode(postCode string) ([]*Address, error)
}
