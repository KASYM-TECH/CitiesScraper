package manager

import (
	"CitiesScratcher/model"
	"sync/atomic"
)

type CurrencyMan struct {
	Currencies map[string]model.Currency
	id         atomic.Int32
}

func NewCurrencyMan() *CurrencyMan {
	return &CurrencyMan{
		Currencies: make(map[string]model.Currency),
	}
}

func (cm *CurrencyMan) Get(currency string) int {
	if c, ok := cm.Currencies[currency]; ok {
		return c.Id
	}
	newCurrency := model.Currency{
		Id:   int(cm.id.Add(1)),
		Name: currency,
	}
	cm.Currencies[currency] = newCurrency

	return newCurrency.Id
}
