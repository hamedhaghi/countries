package countries

import (
	"encoding/json"
)

type Country struct {
	Capital   []string          `json:"capital"`
	Languages map[string]string `json:"languages"`
	Flags     map[string]string `json:"flags"`
	Name      struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
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
	s, err := get("/name/" + name + "?fullText=true")
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

func ByCapital(capital string) (country Country, err error) {
	s, err := get("/capital/" + capital)
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

func ByLanguage(language string) (countries []Country, err error) {
	s, err := get("/lang/" + language)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(s), &countries)
	return
}
