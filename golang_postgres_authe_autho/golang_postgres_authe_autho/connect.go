package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func GetDatabase() *gorm.DB {
	databasename := "userdb"
	database := "postgres"
	databasepassword := "2520"
	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"
	connection, err := gorm.Open(database, databaseurl)
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqldb := connection.DB()

	err = sqldb.Ping()

	if err != nil {
		log.Fatal("database connected")
	}
	fmt.Println("connected to database")
	return connection
}
func InitialMigration() {
	connection := GetDatabase()
	defer Closedatabase(connection)
	connection.AutoMigrate(User{})
}
func Closedatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
