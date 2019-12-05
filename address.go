package jpostcode

import "encoding/json"

type Address struct {
	PostCode       string `json:"postCode"`
	Prefecture     string `json:"prefecture"`
	PrefectureKana string `json:"prefectureKana"`
	PrefectureCode string `json:"prefectureCode"`
	City           string `json:"city"`
	CityKana       string `json:"cityKana"`
	Town           string `json:"town"`
	TownKana       string `json:"townKana"`
	Street         string `json:"street"`
	OfficeName     string `json:"officeName"`
	OfficeNameKana string `json:"officeNameKana"`
}

func (addr *Address) ToJSON() (string, error) {
	b, err := json.Marshal(addr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
