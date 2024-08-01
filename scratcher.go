package main

import (
	"CitiesScratcher/model"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Scratcher struct {
	client *http.Client
}

func NewScratcher() *Scratcher {
	return &Scratcher{client: http.DefaultClient}
}

func (scr *Scratcher) GetCountryCurrencyIso(ctx context.Context) (*model.CountriesCurrencyIsoDto, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://countriesnow.space/api/v0.1/countries/currency", nil)
	if err != nil {
		return nil, err
	}

	resp, err := scr.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	countries := &model.CountriesCurrencyIsoDto{}
	err = json.Unmarshal(body, countries)
	if err != nil {
		return nil, err
	}

	return countries, err
}

func (scr *Scratcher) GetCountriesDialCodes(ctx context.Context) (*model.CountriesCodesDto, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://countriesnow.space/api/v0.1/countries/codes", nil)
	if err != nil {
		return nil, err
	}

	resp, err := scr.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	countries := &model.CountriesCodesDto{}
	err = json.Unmarshal(body, countries)
	if err != nil {
		return nil, err
	}

	return countries, err
}

func (scr *Scratcher) GetCountryCities(ctx context.Context) (*model.CountriesCityDto, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://countriesnow.space/api/v0.1/countries", nil)
	if err != nil {
		return nil, err
	}
	resp, err := scr.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	countryCities := &model.CountriesCityDto{}
	err = json.Unmarshal(body, countryCities)
	if err != nil {
		return nil, err
	}

	return countryCities, err
}
