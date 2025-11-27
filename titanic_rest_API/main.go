package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("titanic.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connectToSQLite()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Perform database migration
	err = db.AutoMigrate(&Tbl_Titanic{})
	if err != nil {
		log.Fatal(err)
	}

	routes()
}
