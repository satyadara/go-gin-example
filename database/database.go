package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"satya-labs/database/models"
)

func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database")
	models.Migrate(db)
	return db, err
}
