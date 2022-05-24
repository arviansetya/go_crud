package config

import (
	"crud/database"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConDB() {
	host := "localhost"
	port := "3306"
	dbname := "crudorm"
	username := "root"
	password := ""

	result := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"

	DB, err = gorm.Open(mysql.Open(result), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}

	for _, model := range database.RegisterModels() {
		err = DB.AutoMigrate(model.Model)

		if err != nil {
			panic("Erorr bro")
		}
	}

	fmt.Println("Connection Succes")
}
