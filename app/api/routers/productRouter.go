// This file contains the routes for the product controller
//
// Usage:
// Use the RegisterProductRouter function to register the product routes to the router
//
// Example:
//
// router := mux.NewRouter()
// routers.RegisterProductRouter(router)
package routers

import (
	"github.com/emso-c/go-assignment/api/controllers"
	"github.com/gorilla/mux"
)

func RegisterProductRouter(router *mux.Router) {
	pc := controllers.NewProductController()

	productRouter := router.PathPrefix("/").Subrouter()
	productRouter.HandleFunc("/happymoons", pc.GetAll).Methods("GET")
	productRouter.HandleFunc("/happymoons/ex={excludedColumns}", pc.GetExcluded).Methods("GET")
	productRouter.HandleFunc("/happymoons/in={includedColumns}", pc.GetIncluded).Methods("GET")
	productRouter.HandleFunc("/happymoons/csv", pc.GetCSV).Methods("GET")
}
