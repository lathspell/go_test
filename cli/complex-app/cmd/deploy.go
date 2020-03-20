package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	log.Print("Init deploy.go")
	cli.AddCommand(&cobra.Command{
		Use:   "deploy",
		Short: "Does something environment specific (see --env or CMPLX_ENV)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Deploying according to Viper in: %s\n", viper.Get("env"))
		},
	})
}
