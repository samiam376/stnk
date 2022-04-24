package avgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFormatRequest(t *testing.T) {
	api_key := "fakekey"
	symbol := "IBM"
	expected := "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=fakekey"
	actual := formatIntraDayRequest(api_key, symbol)

	if expected != actual {
		t.Error("expected did not equal actual")
	}

}

func TestUnmarshallIntraDayPricesToSortedSeries(t *testing.T) {
	data, _ := ioutil.ReadFile("mock_intraday_response.json")
	_, err := unmarshallIntraDayPricesToSortedSeries(data)
	if err != nil {
		t.Errorf("failed to unmarshal with error: %s", err)
	}
}

type mockClient struct{}

func (m mockClient) Get(input string) (*http.Response, error) {
	data, _ := ioutil.ReadFile("mock_response.json")
	r := ioutil.NopCloser(bytes.NewReader(data))
	response := &http.Response{StatusCode: 200, Body: r}
	return response, nil
}

func TestRequestIntraDayPrices(t *testing.T) {
	client := &mockClient{}
	_, err := RequestIntraDayPrices("mockkey", "IBM", client)
	if err != nil {
		t.Error("failed")
	}
}
