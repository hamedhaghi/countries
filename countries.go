package countries

import (
	"encoding/json"
)

type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Country struct {
	Name       Name                `json:"name"`
	Currencies map[string]Currency `json:"currencies"`
}

func All() (countries []Country, err error) {
	s, err := get("/all")
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(s), &countries)
	return
}

func ByName(name string) (country Country, err error) {
	s, err := get("/name/" + name)
	if err != nil {
		return
	}
	countries := []Country{}
	err = json.Unmarshal([]byte(s), &countries)
	if len(countries) > 0 {
		country = countries[0]
	}
	return
}

func ByRegion(region string) (countries []Country, err error) {
	s, err := get("/region/" + region)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(s), &countries)
	return
}
