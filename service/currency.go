package service

import (
	"fmt"
	commona "go-rest-currency-converter/common"
)

type Currency uint8

const (
	TWD Currency = iota
	JPY
	USD
)

var EnumConverter = map[string]Currency{
	"TWD": TWD,
	"JPY": JPY,
	"USD": USD,
}

// TODO: get exchange rate from web or other method
var exchangeRates = map[Currency]map[Currency]float64{
	TWD: {TWD: 1, JPY: 3.669, USD: 0.03281},
	JPY: {TWD: 0.26956, JPY: 1, USD: 0.00885},
	USD: {TWD: 30.444, JPY: 111.801, USD: 1},
}

type CurrencyDTO struct {
	From   Currency
	To     Currency
	Amount float64
}

func ConvertCurrency(params CurrencyDTO) (string, error) {
	rate, ok := exchangeRates[params.From][params.To]
	if !ok {
		return "", fmt.Errorf("unsupported currency conversion: %+v to %+v", params.From, params.To)
	}

	result := params.Amount * rate
	roundedResult := fmt.Sprintf("%.2f", result)
	formattedResult := commona.FormatAmount(roundedResult)

	return formattedResult, nil
}
