package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	log.Print("Init version.go")
	cli.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 1.2.3")
	},
}
