package jpostcode

import "encoding/json"

type Address struct {
	Postcode       string `json:"postcode"`
	Prefecture     string `json:"prefecture"`
	PrefectureKana string `json:"prefecture_kana"`
	PrefectureCode string `json:"prefecture_code"`
	City           string `json:"city"`
	CityKana       string `json:"city_kana"`
	Town           string `json:"town"`
	TownKana       string `json:"town_kana"`
	Street         string `json:"street"`
	OfficeName     string `json:"office_name"`
	OfficeNameKana string `json:"office_name_kana"`
}

func (addr *Address) ToJSON() (string, error) {
	b, err := json.Marshal(addr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
