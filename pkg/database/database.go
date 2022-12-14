package database

import (
	"fmt"

	"github.com/dogab/finance-analyser-api/app/model"
	"github.com/dogab/finance-analyser-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func ConnectDB() {
	var err error

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	// Connect to the DB and initialize the DB variable
	DBconn, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	err = DBconn.AutoMigrate(&model.Entry{})
	if err != nil {
		fmt.Printf("Error migration model Entry: %s", err)
	}
	err = DBconn.AutoMigrate(&model.Category{})
	if err != nil {
		fmt.Printf("Error migration model Category: %s", err)
	}

	fmt.Println("Database Migrated")
}
