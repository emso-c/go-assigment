// This file contains the logic for the product controller
//
// Endpoints:
//
// GET /happymoons
//
// # Returns all products in the database in JSON format
//
// GET /happymoons/ex=category1,category2,...
//
// # Returns all products in the database that match the specified categories in JSON format
//
// GET /happymoons/in=category1,category2,...
//
// # Returns all products in the database that do not match the specified categories in JSON format
//
// GET /happymoons/csv
//
// Returns all products in the database in CSV format
// The CSV file is saved to the local filesystem and
//
// Example usage:
//
// router := mux.NewRouter()
//
// router.HandleFunc("/happymoons", controllers.NewProductController().GetAll).Methods("GET")
//
// router.HandleFunc("/happymoons/ex={excludedColumns}", controllers.NewProductController().GetExcluded).Methods("GET")
//
// router.HandleFunc("/happymoons/in={includedColumns}", controllers.NewProductController().GetIncluded).Methods("GET")
//
// router.HandleFunc("/happymoons/csv", controllers.NewProductController().GetCSV).Methods("GET")
package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/emso-c/go-assignment/api/responses"
	"github.com/emso-c/go-assignment/database"
	"github.com/emso-c/go-assignment/models"
	"github.com/gorilla/mux"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (pc *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	var products []models.GORMProduct
	database.DB.Db.Select("Category, Name, Description, Price").Find(&products)

	formattedProducts := FormatProducts(products)

	err := responses.JSON(w, http.StatusOK, formattedProducts)
	if err != nil {
		panic(err)
	}
}

func (pc *ProductController) GetExcluded(w http.ResponseWriter, r *http.Request) {
	var products []models.GORMProduct
	database.DB.Db.Select("Category, Name, Description, Price").Find(&products)

	formattedProducts := FormatProducts(products)

	// get excluded columns from url
	vars := mux.Vars(r)
	excludedColumns := strings.Split(vars["excludedColumns"], ",")

	// remove excluded columns from json output
	for _, product := range formattedProducts {
		for _, excludedColumn := range excludedColumns {
			delete(product, string(excludedColumn))
		}
	}

	err := responses.JSON(w, http.StatusOK, formattedProducts)
	if err != nil {
		panic(err)
	}
}

func (pc *ProductController) GetIncluded(w http.ResponseWriter, r *http.Request) {
	var products []models.GORMProduct
	database.DB.Db.Select("Category, Name, Description, Price").Find(&products)

	formattedProducts := FormatProducts(products)

	// get included columns from url
	vars := mux.Vars(r)
	includedColumns := strings.Split(vars["includedColumns"], ",")

	// get only included columns from json output
	for _, product := range formattedProducts {
		for key := range product {
			for _, includedColumn := range includedColumns {
				if key != includedColumn {
					delete(product, key)
				}
			}
		}
	}

	err := responses.JSON(w, http.StatusOK, formattedProducts)
	if err != nil {
		panic(err)
	}
}

func (pc *ProductController) GetCSV(w http.ResponseWriter, r *http.Request) {
	var products []models.GORMProduct
	database.DB.Db.Select("Category, Name, Description, Price").Find(&products)

	formattedProducts := FormatProducts(products)

	// modify response headers to download csv file
	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "kategori,urun,aciklama,fiyat\n")
	for _, product := range formattedProducts {
		fmt.Fprintf(w, "%s,%s,\"%s\",%v\n", product["kategori"], product["urun"], product["aciklama"], product["fiyat"])
	}
}

func FormatProducts(products []models.GORMProduct) []map[string]interface{} {
	formattedProducts := make([]map[string]interface{}, len(products))
	for i, product := range products {
		formattedProducts[i] = map[string]interface{}{
			"kategori": product.Category,
			"urun":     product.Name,
			"aciklama": product.Description,
			"fiyat":    product.Price,
		}
	}
	return formattedProducts
}
