package avgo

import (
	"fmt"
	"io"
	"net/http"
)

//price struct defines price at point in time
type Price struct {
	Open   float64 `json:"1. open"`
	High   float64 `json:"2. high"`
	Low    float64 `json:"3. low"`
	Close  float64 `json:"4. close"`
	Volume float64 `json:"5. volume"`
}

//timeseries is a map of a time to a price struct
type Timeseries struct {
	Time map[string]Price `json:"Time Series (5min)"`
}

//response metadata
type Metadata struct {
	Information string `json:"1. Information"`
	Symbol      string `json:"2. Symbol"`
	LastRefresh string `json:"3. Last Refreshed"`
	Interval    string `json:"4. Interval"`
	OutputSize  string `json:"5. Output Size"`
	TimeZone    string `json:"6. Time Zone"`
}

//IntraDay Pricing Api Call response
type IntraDayResponse struct {
	MD Metadata   `json:"Meta Data"`
	TS Timeseries `json:"Time Series (5min)"`
}

//interface to call alpha vantage api
type AvGoClient interface {
	Get(string) (*http.Response, error)
}

//unmarshal api response into Timerseries
func unmarshallIntraDayPricesToTimeseries(body []byte) (*Timeseries, error) {
	return nil, nil
}

//formated query string to include stock symbol and api key
func formatIntraDayRequest(api_key string, symbol string) string {
	request := "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s"
	formatted := fmt.Sprintf(request, symbol, api_key)
	return formatted
}

func RequestIntraDayPrices(api_key string, symbol string, client AvGoClient) (*Timeseries, error) {
	request := formatIntraDayRequest(api_key, symbol)
	response, err := client.Get(request)
	if err != nil {
		return nil, err
	}

	bodyString, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	ts, err := unmarshallIntraDayPricesToTimeseries(bodyString)

	if err != nil {
		return nil, err
	}

	return ts, nil

}
