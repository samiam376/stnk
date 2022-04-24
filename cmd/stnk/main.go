package main

import (
	"fmt"
	"os"

	"github.com/samiam376/stnk/subcmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("expected one of the following subcommands \n intraday \n quote \n")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "intraday":
		subcmd.HandleIntraDay(os.Args[2:])
		os.Exit(0)
	case "quote":
		subcmd.HandleQuote(os.Args[2:])
		os.Exit(0)
	default:
		os.Exit(1)
	}

}
