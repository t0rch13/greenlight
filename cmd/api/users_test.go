package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com.t0rch13.greenlight/internal/data"
	"github.com.t0rch13.greenlight/internal/jsonlog"
)

func TestRegisterUserHandler(t *testing.T) {

	app := &application{
		config: config{},          // Populate with necessary config for testing
		logger: &jsonlog.Logger{}, // Mock logger
		models: &mockModels{},     // Mock data.Models
		mailer: &mockMailer{},     // Mock mailer.Mailer
	}

	// Create a JSON payload for registration
	payload := []byte(`{"name": "Test User", "email": "test@example.com", "password": "password123"}`)

	// Create a request with the payload
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler function
	app.registerUserHandler(rec, req)

	// Check the response status code
	if rec.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d but got %d", http.StatusAccepted, rec.Code)
	}
}

func TestActivateUserHandler(t *testing.T) {
	app := &application{
		config: config{},          // Populate with necessary config for testing
		logger: &jsonlog.Logger{}, // Mock logger
		models: &mockModels{},     // Mock data.Models
		mailer: &mockMailer{},     // Mock mailer.Mailer
	}

	// Create a JSON payload for activation
	payload := []byte(`{"token": "testToken123"}`)

	// Create a request with the payload
	req := httptest.NewRequest("POST", "/activate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to record the response
	rec := httptest.NewRecorder()

	// Call the handler function
	app.activateUserHandler(rec, req)

	// Check the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}
}

// Mock implementations for Models and Mailer interfaces

type mockModels struct{}

func (m *mockModels) Users() data.UserModel {
	return &mockUserModel{}
}

type mockUserModel struct{}

func (m *mockUserModel) Insert(user *data.User) error {
	// Mock Insert implementation
	return nil
}

func (m *mockUserModel) Update(user *data.User) error {
	// Mock Update implementation
	return nil
}

func (m *mockUserModel) GetForToken(scope, token string) (*data.User, error) {
	// Mock GetForToken implementation
	return &data.User{}, nil
}

type mockMailer struct{}

func (m *mockMailer) Send(to, template string, data map[string]interface{}) error {
	// Mock Send implementation
	return nil
}
