package middleware

import (
	"log"               // Import the log package for logging
	"net/http"          // HTTP client and server implementations
	"net/http/httptest" // A package for HTTP testing
	"os"                // Operating system functionality, including file handling
	"testing"           // Support for automated testing of Go packages
)

func TestLoggingMiddleware(t *testing.T) {
	// Setup: Create a new HTTP request to simulate an incoming GET request.
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err) // If creating the request fails, end the test with an error.
	}

	// Create a response recorder to capture the middleware's response.
	rr := httptest.NewRecorder()

	// Create a logger to capture log output in a temporary file.
	logOutput := os.TempDir() + "/log.txt" // Define the log file path.
	// Attempt to open or create the log file.
	file, err := os.OpenFile(logOutput, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal(err) // If opening the log file fails, end the test with an error.
	}
	// Redirect log output to the newly opened file.
	log.SetOutput(file)

	// Define a dummy "next" handler that the middleware will call.
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond with HTTP 200 OK.
	})

	// Apply the Logging middleware to the nextHandler.
	handler := Logging(nextHandler)

	// Serve the HTTP request through the middleware (and next handler).
	handler.ServeHTTP(rr, req)

	// Cleanup: Reset the log output to stderr and close/remove the temporary log file.
	log.SetOutput(os.Stderr) // Reset log output to default.
	file.Close()             // Close the log file.
	os.Remove(logOutput)     // Remove the temporary log file.

	// Check the status code of the response recorded by the response recorder.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Additional checks can be performed here to inspect the log file if needed to ensure
	// the logging occurred as expected. This might involve reading the file and
	// checking its contents against expected log entries.
}
