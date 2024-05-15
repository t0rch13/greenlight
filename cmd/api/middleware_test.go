package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock application
type mockApplication struct{}

func (app *mockApplication) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}

func TestRecoverPanic(t *testing.T) {
	// Create a minimal instance of the application
	app := &application{
		config: config{}, // Initialize with relevant config fields
	}

	// Create a test handler that panics
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	// Wrap the handler with recoverPanic middleware
	wrappedHandler := app.recoverPanic(handler)

	// Create a recorder to record the response
	rec := httptest.NewRecorder()

	// Call the wrapped handler
	wrappedHandler.ServeHTTP(rec, nil)

	// Check if the status code is Internal Server Error
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestRateLimit(t *testing.T) {
	// Create a minimal instance of the application
	app := &application{
		config: config{}, // Initialize with relevant config fields
	}

	// Create a test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the handler with rateLimit middleware
	wrappedHandler := app.rateLimit(handler)

	// Create a recorder to record the response
	rec := httptest.NewRecorder()

	// Call the wrapped handler multiple times within the rate limit
	for i := 0; i < 3; i++ {
		// Call the wrapped handler
		wrappedHandler.ServeHTTP(rec, nil)

		// Check if the status code is OK
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		// Reset the recorder
		rec = httptest.NewRecorder()
	}

	// Call the wrapped handler again, exceeding the rate limit
	wrappedHandler.ServeHTTP(rec, nil)

	// Check if the status code is Too Many Requests
	if rec.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status code %d but got %d", http.StatusTooManyRequests, rec.Code)
	}
}

func TestAuthenticate(t *testing.T) {
	// Create a minimal instance of the application
	app := &application{
		config: config{}, // Initialize with relevant config fields
	}

	// Create a test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the handler with authenticate middleware
	wrappedHandler := app.authenticate(handler)

	// Create a recorder to record the response
	rec := httptest.NewRecorder()

	// Call the wrapped handler without Authorization header
	reqWithoutAuth := httptest.NewRequest("GET", "/", nil)
	wrappedHandler.ServeHTTP(rec, reqWithoutAuth)

	// Check if the status code is OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	// Call the wrapped handler with invalid Authorization header
	reqWithInvalidToken := httptest.NewRequest("GET", "/", nil)
	reqWithInvalidToken.Header.Set("Authorization", "Bearer invalidtoken")
	rec = httptest.NewRecorder()
	wrappedHandler.ServeHTTP(rec, reqWithInvalidToken)

	// Check if the status code is Unauthorized
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, rec.Code)
	}

}
