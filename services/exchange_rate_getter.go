package services

import (
	"encoding/json"
	"io"
	"net/http"
)

type ExchangeRate struct {
	CurrencyCodeA int
	CurrencyCodeB int
	Date          int
	RateBuy       float32
	RateSell      float32
}

func GetExchangeRate() (float32, error) {
	var exchangeRate []ExchangeRate

	url := "https://api.monobank.ua/bank/currency"

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		return 0, err
	}

	return exchangeRate[0].RateBuy, nil
}
