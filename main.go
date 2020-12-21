package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"GSM-Festival-Master-Back/config"
	"GSM-Festival-Master-Back/database"
	"GSM-Festival-Master-Back/rest"
)

func main() {
	config.InitConfig()
	database.Connect()
	rest.RunAPI(":1324")
}
