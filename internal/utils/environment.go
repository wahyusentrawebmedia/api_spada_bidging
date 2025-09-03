package utils

import (
	"api/spada/internal/database"
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Set up Viper to read from .env file
	viper.SetConfigFile(".env") // specify the .env file directly
	viper.SetConfigType("env")
	viper.AddConfigPath(".") // look for .env in the working directory

	// Read in environment variables as well
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf(".env file not found: %v, relying on environment variables", err)
	}

	// Initialize database connection
	err = database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}
