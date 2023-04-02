package models

// Represents a scraper
type Scraper interface {
	ScrapeData() ([]Product, error)
}
