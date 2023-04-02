// Author: Emir Soyalan - emirsoyalan79@gmail.com
//
// License: MIT
//
// This project is aimed to provide a simple and easy to use data storage API.
// It is written in Go and uses PostgreSQL as the database.
// The API is designed to be RESTful and ready to be consumed by any client.
// The data is scraped from the Happy Moons website.
//
// This file is the entry point for the application. It is responsible for
// initializing the API server and loading any necessary application-level
// configurations.
//
// For more documentation, see the following URL after running `godoc -http=127.0.0.1:6060`
// http://localhost:6060/pkg/github.com/emso-c/go-assignment/
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/emso-c/go-assignment/api"
	"github.com/emso-c/go-assignment/config"
	"github.com/emso-c/go-assignment/database"
	"github.com/emso-c/go-assignment/models"
	"github.com/emso-c/go-assignment/scraper"
	"github.com/emso-c/go-assignment/utils/console"
)

func main() {
	// Configure log flags
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate | log.Lmicroseconds)

	// Load configuration values
	cErr := config.LoadEnv(config.DEFAULT_CONFIG_PATH)
	if cErr != nil {
		log.Fatal(cErr)
		return
	}

	// Initialize the database
	database.ConnectDb()
	// If the database is empty, scrape data and insert into database
	if database.DB.Db.Find(&models.GORMProduct{}).RowsAffected == 0 {
		log.Print("Database is empty, scraping data...")
		products, err := scraper.NewHappyMoonScraper(true).ScrapeData()
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Scraping complete, inserting data into database...")
		for _, product := range products {
			database.DB.Db.Create(&models.GORMProduct{Product: product})
		}
		log.Print("Insertion complete")
	}

	// Initialize the API server
	api.Init()
	router := api.GetRouter()

	// Start the API server
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	log.Print(
		"Starting server on ",
		console.RED+console.UNDERLINE,
		"http://"+host+":"+port,
		console.RESET,
	)
	hErr := http.ListenAndServe(":"+port, router)
	if hErr != nil {
		log.Fatal(hErr)
	}
}
