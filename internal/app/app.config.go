package app

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

const (
	host     = "localhost"   // localport
	port     = 5432 // (port for localhost)
	user     = "postgres"   // postgres user
	password = ""   // replace with your postgres password
	dbname   = ""   // replace with your postgres DB name
)

var err error

func SetupDB() {
	// postgres localhost connection
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})  // gorm init
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&Goly{})     // create the table
	if err != nil {
		log.Print(err.Error())
	}
}