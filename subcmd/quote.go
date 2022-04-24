package subcmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samiam376/stnk/pkg/avgo"
)

func HandleQuote(args []string) {
	quote := flag.NewFlagSet("quote", flag.ExitOnError)
	apiKey := quote.String("api-key", "", "manually provide alpha vantage api key")
	symbol := quote.String("symbol", "", "stock symbol")
	quote.Parse(args)

	if *apiKey == "" {
		fmt.Println("Expected api key from alpha vantage")
		os.Exit(1)
	}

	if *symbol == "" {
		fmt.Println("Expected stock symbol")
		os.Exit(1)
	}

	client := &http.Client{}
	q, err := avgo.RequestQuote(*apiKey, *symbol, client)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	avgo.PrintQuote(q)
}
