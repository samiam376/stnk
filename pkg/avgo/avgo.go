package avgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

//price struct defines price at point in time
type Price struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume float64 `json:"5. volume,string"`
}

//timeseries is a map of a time to a price struct
type Timeseries map[string]Price

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
// Timeseries is not neccisarily sorted
type IntraDayResponse struct {
	MD Metadata   `json:"Meta Data"`
	TS Timeseries `json:"Time Series (60min)"`
}

//Read only container that holds the sorted keys to the timeseries and the underlying map
//these can be used to iterate over the container in order
type SortedSeries struct {
	unsortedContainer Timeseries
	SortedKeys        []string
}

func NewSortedSeries(unsorted Timeseries) *SortedSeries {
	//copy over unsorted map to prevent mutations on the underlying data
	container := Timeseries{}
	sortedKeys := make([]string, 0)
	for key, val := range unsorted {
		container[key] = val
		sortedKeys = append(sortedKeys, key)
	}
	//sort the keys
	sort.Strings(sortedKeys)
	//return a new object
	return &SortedSeries{unsortedContainer: container, SortedKeys: sortedKeys}
}

func (s *SortedSeries) Get(key string) Price {
	return s.unsortedContainer[key]
}

func PrintTailAsc(ss *SortedSeries, n *int) {
	startIdx := 0
	l := len(ss.SortedKeys)
	if n != nil {
		startIdx = l - *n - 1
		if 0 > startIdx {
			startIdx = 0
		}
	}

	divider := strings.Repeat("-", 15)
	for i := startIdx; i < l; i++ {
		time := ss.SortedKeys[i]
		price := ss.Get(time)
		fmt.Println(divider)
		fmt.Printf("Time: %s/n", time)
		fmt.Printf("Open: %f/n", price.Open)
		fmt.Printf("Close: %f/n", price.Close)
		fmt.Printf("Low: %f/n", price.Low)
		fmt.Printf("High: %f/n", price.High)
		fmt.Println(divider)
	}
}

//interface to call alpha vantage api
type AvGoClient interface {
	Get(string) (*http.Response, error)
}

//unmarshal api response into Timerseries
func unmarshallIntraDayPricesToSortedSeries(body []byte) (*SortedSeries, error) {
	resp := &IntraDayResponse{}
	err := json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	ss := NewSortedSeries(resp.TS)
	return ss, nil
}

//formated query string to include stock symbol and api key
func formatIntraDayRequest(api_key string, symbol string) string {
	request := "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s"
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
