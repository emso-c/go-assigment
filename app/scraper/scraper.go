// Package scraper contains the implementation of the Scraper interface
//
// The scraper is responsible for fetching data from the specified source and
// parsing it into a Product array.
//
// Example usage:
//
//	scraper := NewHappyMoonScraper(false)
//	doc, err := scraper.InitScraper()
//	if err != nil {
//		panic(err)
//	}
//	products := scraper.ParseData(doc)
//	fmt.Println(products)
package scraper

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/emso-c/go-assignment/models"
)

// Scraper for the Happy Moon website
type HappyMoonScraper struct {
	url    string
	silent bool
}

// Creates a new HappyMoonScraper instance
//
// silent: if true, the scraper will not log any messages
//
// returns: a new HappyMoonScraper instance
func NewHappyMoonScraper(silent bool) models.Scraper {
	return &HappyMoonScraper{
		url:    os.Getenv("SCRAPER_URL"),
		silent: silent,
	}

}

// Initializes the scraper by fetching the HTML document from the specified URL
// and parsing it into a goquery document
//
// returns: a goquery document
func (s *HappyMoonScraper) InitScraper() (*goquery.Document, error) {
	if !s.silent {
		log.Println("Scraping data from", s.url)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Parses the goquery document and extracts the data into a Product array
//
// doc: the goquery document
//
// returns: a Product array
func (s *HappyMoonScraper) ParseData(doc *goquery.Document) []models.Product {
	currentCategory := ""
	products := []models.Product{}

	// Find the first ul element with class "menu"
	menu := doc.Find("ul.menu").First()
	// Iterate through each li element within the ul element
	menu.Find("li").Each(func(i int, qs *goquery.Selection) {
		// For each li element, check class name
		class, _ := qs.Attr("class")
		if class == "menuCategory" {
			currentCategory = qs.Find("div.categoryName").Text()
			currentCategory = strings.TrimSpace(currentCategory)
			return
		}
		// if class is "product bubbleImageSideAligned", then extract data
		if class == "product bubbleImageSideAligned" {
			name := qs.Find("div.itemName").Text()
			name = strings.TrimSpace(name)
			description := qs.Find("div.itemDescription").Text()
			description = strings.TrimSpace(description)
			price := qs.Find("div.itemPrice").Text()
			price = strings.TrimSpace(price)
			price = price[:len(price)-7]
			priceUint16, err := strconv.ParseUint(price, 10, 16)
			if err != nil {
				log.Fatal(err)
			}
			// Create new Product and append to products array
			products = append(products, models.Product{
				Name:        name,
				Description: description,
				Price:       uint16(priceUint16),
				Category:    currentCategory,
			})
			if !s.silent {
				fmt.Println("Scraped data: -----------------")
				fmt.Println("Name:", name)
				fmt.Println("Description:", description)
				fmt.Println("Price:", price)
				fmt.Println("Category:", currentCategory)
			}
		}
	})
	return products
}

// Scrapes data from the specified URL and parses it into a Product array
//
// returns: a Product array
func (s *HappyMoonScraper) ScrapeData() ([]models.Product, error) {
	doc, err := s.InitScraper()
	if err != nil {
		return nil, err
	}

	products := s.ParseData(doc)

	if !s.silent {
		log.Println("Scraping finished!", len(products), "products found.")
	}
	return products, nil
}
