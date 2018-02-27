package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/olshevskiy87/gopuny/punyurl"
)

const (
	actionShort  = "short"
	actionExpand = "expand"
)

func main() {
	var args struct {
		Action string `arg:"positional,required" help:"\"short\" or \"expand\""`
		URL    string `arg:"positional,required" help:"url to short or expand"`
	}
	argParser := arg.MustParse(&args)

	puny, err := punyurl.New(args.URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch args.Action {
	case actionShort:
		res, err := puny.Short()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(res)
	case actionExpand:
		res, err := puny.Expand()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(res.URL)
	default:
		fmt.Fprintf(os.Stderr, "invalid action: %s\n", args.Action)
		argParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
