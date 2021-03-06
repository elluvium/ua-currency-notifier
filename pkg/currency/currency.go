package currency

import (
	"currency-notifier/models"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func getCurrencies(url string) ([]byte, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func jsonParser() (*models.CurrencyState, error) {
	var state models.CurrencyState

	response, err := getCurrencies("http://resources.finance.ua/ua/public/currency-cash.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(response), &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func StateInitializer() (*models.CurrencyState, error) {
	currencies, err := jsonParser()
	if err != nil {
		return nil, err
	}
	return currencies, nil
}
