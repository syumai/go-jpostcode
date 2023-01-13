package address

import (
	"github.com/mitchellh/mapstructure"
	"github.com/syumai/go-jpostcode"
)

func FromMap(input interface{}) (*jpostcode.Address, error) {
	var addr jpostcode.Address
	if err := mapstructure.Decode(input, &addr); err != nil {
		return nil, err
	}
	return &addr, nil
}
