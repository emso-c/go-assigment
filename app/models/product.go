package models

import "gorm.io/gorm"

// Product represents a product in the database
// It is used to store the data scraped from the website
type Product struct {
	Category    string
	Name        string
	Description string
	Price       uint16
}

// GORMProduct is a wrapper for the Product struct
// for use with the GORM library
type GORMProduct struct {
	gorm.Model
	Product
}
