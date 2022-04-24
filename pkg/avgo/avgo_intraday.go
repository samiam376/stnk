package avgo

import (
	"fmt"
	"io"
)

//formated query string to include stock symbol and api key
func formatIntraDayRequest(api_key string, symbol string) string {
	request := "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=60min&apikey=%s"
	formatted := fmt.Sprintf(request, symbol, api_key)
	return formatted
}

func RequestIntraDayPrices(api_key string, symbol string, client AvGoClient) (*SortedSeries, error) {
	request := formatIntraDayRequest(api_key, symbol)
	response, err := client.Get(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("API Call failed with status code %d", response.StatusCode)
	}

	bodyString, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	ss, err := unmarshallIntraDayPricesToSortedSeries(bodyString)

	if err != nil {
		return nil, err
	}

	return ss, nil

}
