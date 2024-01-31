package main

import (
	_ "github.com/spf13/cobra"
	"github.com/staketab/go-mina-payout/internal"
	"github.com/staketab/go-mina-payout/internal/vars"
	"os"
)

func main() {
	if err := internal.RootCmd.Execute(); err != nil {
		vars.ErrorLog.Println(err)
		os.Exit(1)
	}
}
