package model

type Country struct {
	Id         int    `csv:"id"`
	Name       string `csv:"name"`
	Code       string `csv:"code"`
	Iso2       string `csv:"iso2"`
	Iso3       string `csv:"iso3"`
	CurrencyId int    `csv:"currencyId"`
}

type City struct {
	Id        int    `csv:"id"`
	Name      string `csv:"name"`
	CountryId int    `csv:"countryId"`
}

type Currency struct {
	Id   int    `csv:"id"`
	Name string `csv:"name"`
	Code string `csv:"code"`
}
