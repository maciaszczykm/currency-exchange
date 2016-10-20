package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// apiUrl with exchange rates based on specified currency.
const apiUrl = "https://api.fixer.io/latest?base=%s"

// GetCurrencyRates returns map containing pairs of currencies and their actual exchange rates.
func GetCurrencyRates(base string) (map[string]float64, error) {
	url := fmt.Sprintf(apiUrl, base)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot acess %s", url)
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var apiObject map[string]*json.RawMessage
	err = json.Unmarshal(body, &apiObject)
	if err != nil {
		return nil, err
	}

	var rates map[string]float64
	err = json.Unmarshal(*apiObject["rates"], &rates)
	if err != nil {
		return nil, err
	}

	return rates, nil
}
