package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging middleware logs the HTTP request method, URI, and execution time
func Logging(next http.Handler) http.Handler {
	// Return a http.Handler function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the current time
		start := time.Now()
		// Log the start of the request including method and URI
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
		// Calculate and log the time taken to execute the request
		log.Printf("Completed in %v", time.Since(start))
	})
}
