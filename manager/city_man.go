package manager

import (
	"CitiesScratcher/model"
	"sync/atomic"
)

type CityMan struct {
	Cities []*model.City
	id     atomic.Int32
}

func NewCityMan() *CityMan {
	return &CityMan{}
}

func (cm *CityMan) Save(cityName string, countryId int) {
	newCity := model.City{
		Id:        int(cm.id.Add(1)),
		Name:      cityName,
		CountryId: countryId,
	}
	cm.Cities = append(cm.Cities, &newCity)
}
