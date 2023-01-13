package jpostcode

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type Address struct {
	PostCode       string `json:"postCode" mapstructure:"postcode"`
	Prefecture     string `json:"prefecture" mapstructure:"prefecture"`
	PrefectureKana string `json:"prefectureKana" mapstructure:"prefecture_kana"`
	PrefectureCode int    `json:"prefectureCode" mapstructure:"prefecture_code"`
	City           string `json:"city" mapstructure:"city"`
	CityKana       string `json:"cityKana" mapstructure:"city_kana"`
	Town           string `json:"town" mapstructure:"town"`
	TownKana       string `json:"townKana" mapstructure:"town_kana"`
	Street         string `json:"street" mapstructure:"street"`
	OfficeName     string `json:"officeName" mapstructure:"office_name"`
	OfficeNameKana string `json:"officeNameKana" mapstructure:"office_name_kana"`
}

func (addr *Address) ToJSON() (string, error) {
	b, err := json.Marshal(addr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func AddressFromMap(input interface{}) (*Address, error) {
	var addr Address
	if err := mapstructure.Decode(input, &addr); err != nil {
		return nil, err
	}
	return &addr, nil
}
