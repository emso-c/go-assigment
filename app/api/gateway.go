// Package api provides the HTTP API implementation for the application.
//
// The package initializes the router and registers all the necessary middlewares
// to handle CORS, preflight requests, 404 errors, and rate limiting. Currently, the
// authentication and authorization middlewares are not implemented since the
// assignment does not require endpoints to be protected.
//
// Usage:
// Call the Init function to initialize the router and register the necessary middlewares.
// Use the GetRouter function to retrieve the initialized router.
//
// Example:
//
// // Initialize the router and register the necessary middlewares
// api.Init()
//
// // Retrieve the initialized router
// router := api.GetRouter()
//
// // Start the server and listen for incoming requests
// log.Fatal(http.ListenAndServe(":8080", router))
//
// Notes:
// This package relies on the middlewares and utils/limiter packages for middleware functions and
// rate limiting implementation respectively. Make sure those packages are imported before using this package.
//
// This package also depends on the routers package for registering endpoint handlers. Make sure to
// register all the necessary routers in the Init function before using this package.
package api

import (
	"github.com/emso-c/go-assignment/api/middlewares"
	"github.com/emso-c/go-assignment/api/routers"
	"github.com/emso-c/go-assignment/utils/limiter"
	"github.com/gorilla/mux"
)

// router is the instance of the router that is initialized by the Init function.
var router *mux.Router

// Init initializes the router and registers all the necessary middlewares
// to handle CORS, preflight requests, 404 errors, and rate limiting. Currently, the
// authentication and authorization middlewares are not implemented since the
// assignment does not require endpoints to be protected.
func Init() {
	router = mux.NewRouter()
	// Register routers
	routers.RegisterProductRouter(router)

	// TODO: Add authentication & authorization middlewares.
	// Currently not implemented since the assignment does
	// not require endpoints to be protected.

	// Handle CORS and preflight middleware
	router.Use(middlewares.CORSMiddleware())
	// Handle 404 middleware
	router.NotFoundHandler = middlewares.NotFoundMiddleware()
	// Handle rate limit middleware
	limiter.GetLimiter().Initialize()
	router.Use(middlewares.RateLimitMiddleware())
}

// GetRouter retrieves the instance of the initialized router.
func GetRouter() *mux.Router {
	return router
}
