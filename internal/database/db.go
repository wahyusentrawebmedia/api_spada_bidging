package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASS")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	name := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, pass)
	fmt.Println("Connecting to database with DSN:", dsn) // Debugging line
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
