package main

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	log.Print("Init Viper")

	// Configuration Management
	viper.Set("Verbose", true)

	// Look in environment for "CMPLX_*" variables
	viper.SetEnvPrefix("cmplx")
	viper.AutomaticEnv()

	// Look for complex-app.yml
	viper.SetConfigName("complex-app") // base name of config file
	viper.SetConfigType("yml")   // REQUIRED
	viper.AddConfigPath(".")

	log.Print("Loading complex-app.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Could not read complex-app.yml: %s\n", err)
	}
}
