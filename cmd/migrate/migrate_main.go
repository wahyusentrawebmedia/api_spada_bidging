package main

import (
	"api/spada/internal/database"
	"api/spada/internal/utils"
	"log"
)

func main() {
	utils.InitConfig()
	database.Migrate(database.DB)
	log.Println("Migration completed successfully.")
}
