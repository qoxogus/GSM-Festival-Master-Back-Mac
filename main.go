package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/qoxogus/GSM-Festival-Master-Back/config"
	"github.com/qoxogus/GSM-Festival-Master-Back/database"
	"github.com/qoxogus/GSM-Festival-Master-Back/rest"
)

func main() {
	config.InitConfig()
	database.Connect()
	rest.RunAPI(":1324")
}
