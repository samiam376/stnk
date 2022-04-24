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
	last := intraday.Int("last", 5, "Shows the last n prices in 5 min durations, to show all prices put -1, defaults to 5")
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
	fmt.Printf("Last %d prices for %s\n", last, *symbol)
	avgo.PrintTailAsc(ss, last)
}
