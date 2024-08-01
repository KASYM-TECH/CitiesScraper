package model

// CountryCodesDto countries dial codes
type CountryCodesDto struct {
	Name     string `json:"name"`
	DialCode string `json:"dial_code"`
}

type CountriesCodesDto struct {
	Error   bool              `json:"error"`
	Message string            `json:"msg"`
	Data    []CountryCodesDto `json:"data"`
}

// CountryCurrencyIsoDto currencies and iso
type CountryCurrencyIsoDto struct {
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Iso2     string `json:"iso2"`
	Iso3     string `json:"iso3"`
}

type CountriesCurrencyIsoDto struct {
	Error   bool                    `json:"error"`
	Message string                  `json:"msg"`
	Data    []CountryCurrencyIsoDto `json:"data"`
}

// CountryCityDto country cities
type CountryCityDto struct {
	Country string   `json:"country"`
	Cities  []string `json:"cities"`
}

type CountriesCityDto struct {
	Error   bool             `json:"error"`
	Message string           `json:"msg"`
	Data    []CountryCityDto `json:"data"`
}
