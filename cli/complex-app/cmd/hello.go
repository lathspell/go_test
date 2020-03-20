package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	log.Print("Init hello.go")
	cli.AddCommand(&cobra.Command{
		Use:   "hello",
		Short: "Prints hello <args>",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello %s\n", strings.Join(args, ", "))
		},
	})
}
