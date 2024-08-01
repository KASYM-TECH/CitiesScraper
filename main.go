package main

//f, err := os.Create("trace.out")
//if err != nil {
//	panic(err)
//}
//defer f.Close()
//err = trace.Start(f)
//defer trace.Stop()
import (
	"CitiesScratcher/manager"
	"CitiesScratcher/model"
	"cmp"
	"context"
	"github.com/gocarina/gocsv"
	"golang.org/x/sync/errgroup"
	"os"
	"slices"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	scratcher := NewScratcher()

	var countryCurrDto *model.CountriesCurrencyIsoDto
	var countryCodesDto *model.CountriesCodesDto
	var countryCitiesDto *model.CountriesCityDto

	eg.Go(func() error {
		var err error
		countryCurrDto, err = scratcher.GetCountryCurrencyIso(ctx)
		return err
	})
	eg.Go(func() error {
		var err error
		countryCodesDto, err = scratcher.GetCountriesDialCodes(ctx)
		return err
	})
	eg.Go(func() error {
		var err error
		countryCitiesDto, err = scratcher.GetCountryCities(ctx)
		return err
	})

	err := eg.Wait()
	handleError(err)

	countriesMap := make(map[string]*model.Country, len(countryCurrDto.Data))
	currencyMan := manager.NewCurrencyMan()
	cityMan := manager.NewCityMan()

	for index, countryDto := range countryCurrDto.Data {
		currencyId := currencyMan.Get(countryDto.Currency)
		countriesMap[countryDto.Name] = &model.Country{
			Id:         index + 1,
			Name:       countryDto.Name,
			Iso2:       countryDto.Iso2,
			Iso3:       countryDto.Iso3,
			CurrencyId: currencyId,
		}
	}

	for _, countryDto := range countryCodesDto.Data {
		if country, ok := countriesMap[countryDto.Name]; ok {
			country.Code = countryDto.DialCode
			continue
		}
		println("did not find country name for: " + countryDto.Name)
	}

	for _, countryCities := range countryCitiesDto.Data {
		country, ok := countriesMap[countryCities.Country]
		if !ok {
			panic("did not find country")
		}
		for _, cityName := range countryCities.Cities {
			cityMan.Save(cityName, country.Id)
		}
	}

	egSave, ctx := errgroup.WithContext(context.Background())

	egSave.Go(func() error {
		countriesCsv, err := os.Create("countries.csv")
		defer countriesCsv.Close()
		if err != nil {
			return err
		}
		var countriesArr []*model.Country
		for _, country := range countriesMap {
			countriesArr = append(countriesArr, country)
		}
		slices.SortFunc(countriesArr, func(a, b *model.Country) int {
			if n := cmp.Compare(a.Id, b.Id); n != 0 {
				return n
			}
			return cmp.Compare(a.Name, b.Name)
		})
		return gocsv.MarshalFile(&countriesArr, countriesCsv)
	})

	egSave.Go(func() error {
		CurrenciesCsv, err := os.Create("currencies.csv")
		defer CurrenciesCsv.Close()
		if err != nil {
			return err
		}
		currencyMap := currencyMan.Currencies
		var currencyArr []*model.Currency
		for _, currency := range currencyMap {
			currencyArr = append(currencyArr, &currency)
		}
		slices.SortFunc(currencyArr, func(a, b *model.Currency) int {
			if n := cmp.Compare(a.Id, b.Id); n != 0 {
				return n
			}
			return cmp.Compare(a.Name, b.Name)
		})
		return gocsv.MarshalFile(&currencyArr, CurrenciesCsv)
	})

	egSave.Go(func() error {
		citiesCsv, err := os.Create("cities.csv")
		defer citiesCsv.Close()
		if err != nil {
			return err
		}
		return gocsv.MarshalFile(&cityMan.Cities, citiesCsv)
	})

	err = egSave.Wait()
	handleError(err)
}
