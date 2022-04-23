package main

import (
	"fmt"
	"os"

	"github.com/samiam376/stnk/subcmd"
)

func main() {

	switch os.Args[1] {
	case "intraday":
		subcmd.HandleIntraDay(os.Args[2:])
	default:
		fmt.Print("expected subcommand")
		os.Exit(1)
	}

}
