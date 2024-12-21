package initializers

import (
	"awesomeProject12/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	_ "github.com/go-sql-driver/mysql" // SQL driver for raw database connection
)

// an variable  that we can use outside of this file
var DB *gorm.DB

// func that will let us allow to create or connect to the database
func ConnectDB() {
	//an var error
	var err error
	//all var from .env file for db
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	//connect to db
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	//open db
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if there is error
	if err != nil {
		// it will type out the error
		panic("Failed to connect to database: " + err.Error())
	}
}

// func that will migrate tables in the database
func CreateDb() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to create tables: " + err.Error())
	}
}
