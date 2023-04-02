// Package middlewares provides HTTP middlewares for the API.
//
// Usage:
// Use the CORSMiddleware function as a middleware in your HTTP handlers to enable CORS.
// The middleware retrieves the origin from the request header, and sets the Access-Control-Allow-Origin
// header in the response if the origin is allowed. It also sets the Access-Control-Allow-Methods header
// in the response for a Preflighted OPTIONS request.
//
// Example:
//
// http.Handle("/api/products", middlewares.CORSMiddleware()(http.HandlerFunc(handler)))
//
// The middleware returns a 400 Bad Request error if the remote address is missing, and a 429 Too Many
// Requests error if the remote address has exceeded the rate limit.
package middlewares

import (
	"net/http"
	"os"
	"strings"
)

// CORSMiddleware returns a middleware that enables CORS.
func CORSMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowedOrigins := strings.Split(",", os.Getenv("HTTP_ALLOWED_ORIGINS"))
			if origin != "" {
				for _, allowedOrigin := range allowedOrigins {
					if allowedOrigin == origin {
						w.Header().Set("Access-Control-Allow-Origin", origin)
						break
					}
				}
			}

			// Stop here for a Preflighted OPTIONS request
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Methods", os.Getenv("HTTP_ALLOWED_METHODS"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
