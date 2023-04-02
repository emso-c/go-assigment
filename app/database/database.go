// Package database provides a GORM-based database connection and migration functionality.
//
// Usage:
// Before using the database, set the environment variables DB_USER, DB_PASSWORD, and
// DB_NAME to the corresponding values of your PostgreSQL database. Then initialize the
// database connection and migration by calling the ConnectDb function once at the start of
// your application.
//
// Example:
// Set the environment variables:
// export DB_USER=your-db-username
// export DB_PASSWORD=your-db-password
// export DB_NAME=your-db-name
//
// Initialize the database connection and migration:
// database.ConnectDb()
//
// Now you can use the DB.Db instance to interact with your database. For example:
//
// // Query all products
// var products []models.GORMProduct
// database.DB.Db.Find(&products)
//
// // Create a new product
// product := models.GORMProduct{Name: "New Product", Price: 9.99}
// database.DB.Db.Create(&product)
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/emso-c/go-assignment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBInstance represents a GORM-based database connection instance.
type DBInstance struct {
	Db *gorm.DB
}

// DB is the shared instance of the DBInstance.
var DB DBInstance

// ConnectDb connects to the PostgreSQL database using the environment variables DB_USER,
// DB_PASSWORD, and DB_NAME, and initializes the migration of the database.
func ConnectDb() {
	log.Println("Connecting to database...")
	dsn := fmt.Sprintf(
		"host=database user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db_tmp, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected to database.")
	db_tmp.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating database...")
	db_tmp.AutoMigrate(&models.GORMProduct{})

	log.Println("Database migration complete.")
	DB = DBInstance{Db: db_tmp}
}
