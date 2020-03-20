package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cli = &cobra.Command{Use: "complex"}

func CliHandling() {
	log.Print("Entering CLI Handling")
	if err := cli.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}
