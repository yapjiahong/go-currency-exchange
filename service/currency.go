package service

import (
	"fmt"
	"github.com/shopspring/decimal"
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

func isSupportedCurrency(from, to Currency) (float64, error) {
	rate, ok := exchangeRates[from][to]
	if !ok {
		return 0, fmt.Errorf("unsupported currency conversion: %+v to %+v", from, to)
	}

	return rate, nil
}

func ConvertCurrency(params CurrencyDTO) (string, error) {
	rate, err := isSupportedCurrency(params.From, params.To)
	if err != nil {
		return "", err
	}

	result := params.Amount * rate
	if result > 0 {
		roundedResult, goErr := decimal.NewFromFloat(result).Round(2).Float64()
		if goErr == true {
			return "", fmt.Errorf("float Round error, input: %+v", result)
		}

		result = roundedResult
	}

	roundedResultString := fmt.Sprintf("%.2f", result)
	formattedResult := commona.FormatAmount(roundedResultString)

	return formattedResult, nil
}
