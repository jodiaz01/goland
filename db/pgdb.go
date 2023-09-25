package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=postgres password=admin dbname=db23go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func DBConexion() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Conected to db23go")
	}
}
