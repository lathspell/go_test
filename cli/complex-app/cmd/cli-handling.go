package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cli = &cobra.Command{Use: "complex"}

var Verbose bool

// Command line parameter, use `viper.Get("env")` for actual work!
var Env string

func CliHandling() {
	log.Print("Entering CLI Handling")

	cli.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// The Viper variable "env" can be overridden by the "--env=foo" command line parameter
	cli.PersistentFlags().StringVar(&Env, "env", "dev", "Environment selection")
	viper.BindPFlag("env", cli.PersistentFlags().Lookup("env"))

	if err := cli.Execute(); err != nil {
		fmt.Printf("\nError: %s\n", err.Error())
		os.Exit(1)
	}
}
