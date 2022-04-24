package avgo

import (
	"net/http"
)

//interface to call alpha vantage api
type AvGoClient interface {
	Get(string) (*http.Response, error)
}
