package avgo

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func PrintTailAsc(ss *SortedSeries, n *int) {
	startIdx := 0
	l := len(ss.SortedKeys)
	// if n is provided take last n elements
	if n != nil {
		if *n == -1 {
			startIdx = 0
		} else {
			startIdx = l - *n - 1
			//check bounds
			if 0 > startIdx {
				startIdx = 0
			}
		}
	}

	divider := strings.Repeat("-", 15)
	for i := startIdx; i < l; i++ {
		time := ss.SortedKeys[i]
		price := ss.Get(time)
		fmt.Println(divider)
		fmt.Printf("Time: %s\n", time)
		fmt.Printf("Open: %f\n", price.Open)
		fmt.Printf("Close: %f\n", price.Close)
		fmt.Printf("Low: %f\n", price.Low)
		fmt.Printf("High: %f\n", price.High)
		fmt.Println(divider)
	}
}
