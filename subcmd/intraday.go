package subcmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samiam376/stnk/pkg/avgo"
)

func HandleIntraDay(args []string) {
	intraday := flag.NewFlagSet("intraday", flag.ExitOnError)
	apiKey := intraday.String("api-key", "", "manually provide alpha vantage api key")
	symbol := intraday.String("symbol", "", "stock symbol")

	intraday.Parse(args)
	if *apiKey == "" {
		fmt.Println("Expected api key from alpha vantage")
		os.Exit(1)
	}

	if *symbol == "" {
		fmt.Println("Expected stock symbol")
		os.Exit(1)
	}

	client := &http.Client{}
	ss, err := avgo.RequestIntraDayPrices(*apiKey, *symbol, client)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	take := 5
	avgo.PrintTailAsc(ss, &take)
}
