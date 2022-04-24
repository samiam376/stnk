package avgo

import (
	"encoding/json"
	"fmt"
	"io"
)

func formatQuoteRequest(api_key string, symbol string) string {
	request := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s"
	formatted := fmt.Sprintf(request, symbol, api_key)
	return formatted
}

func RequestQuote(api_key string, symbol string, client AvGoClient) (*Quote, error) {
	request := formatQuoteRequest(api_key, symbol)
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

	resp := &Quote{}
	err = json.Unmarshal(bodyString, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func PrintQuote(quote *Quote) {
	fmt.Printf("Symbol: %s\n", quote.Symbol)
	fmt.Printf("Open: %f\n", quote.Open)
	fmt.Printf("High: %f\n", quote.High)
	fmt.Printf("Low: %f\n", quote.Low)
	fmt.Printf("Price: %f\n", quote.Price)
	fmt.Printf("Volume: %f\n", quote.Volume)
	fmt.Printf("Latest Trading Day: %s\n", quote.LatestTradingDay)
	fmt.Printf("Change: %f\n", quote.Change)
	fmt.Printf("Change Pct: %s\n", quote.ChangePct)
}
