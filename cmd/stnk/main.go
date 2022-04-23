package main

import (
	"fmt"
	"os"

	"github.com/samiam376/stnk/subcmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("expected subcommand see --help \n")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "intraday":
		subcmd.HandleIntraDay(os.Args[2:])
	default:
		os.Exit(1)
	}

}
