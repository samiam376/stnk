package avgo

import "sort"

// defines a singe quote
type Quote struct {
	Symbol           string  `json:"01. symbol:"`
	Open             float64 `json:"02. open:,string"`
	High             float64 `json:"03. high:,string"`
	Low              float64 `json:"04: low:,string"`
	Price            float64 `json:"05. price:,string"`
	Volume           float64 `json:"06. volume:,string"`
	LatestTradingDay string  `json:"07. latest trading day:"`
	PreviousClose    float64 `json:"08. previous close:,string"`
	Change           float64 `json:"09. change:,string"`
	ChangePct        string  `json:"10. change percent:"`
}

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
