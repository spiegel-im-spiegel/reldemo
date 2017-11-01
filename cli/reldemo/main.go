package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/reldemo/cli/reldemo/facade"
)

func main() {
	os.Exit(facade.Execute(gocli.NewUI(
		gocli.Reader(os.Stdin),
		gocli.Writer(os.Stdout),
		gocli.ErrorWriter(os.Stderr),
	)).Int())
}
