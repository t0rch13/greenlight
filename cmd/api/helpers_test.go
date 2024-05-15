package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestReadIDParam_ValidID(t *testing.T) {
	app := &application{}

	// Create a mock request with a valid ID parameter
	req := &http.Request{
		Header: http.Header{},
	}
	params := httprouter.Params{
		{Key: "id", Value: "123"},
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
	req = req.WithContext(ctx)

	// Call the helper function
	id, err := app.readIDParam(req)

	// Check if the returned ID is correct
	if id != 123 {
		t.Errorf("Expected ID 123 but got %d", id)
	}

	// Check if there's no error
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}
}

func TestReadIDParam_InvalidID(t *testing.T) {
	app := &application{}

	// Create a mock request with an invalid ID parameter
	req := &http.Request{
		Header: http.Header{},
	}
	params := httprouter.Params{
		{Key: "id", Value: "invalid"},
	}
	ctx := req.Context()
	ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
	req = req.WithContext(ctx)

	// Call the helper function
	id, err := app.readIDParam(req)

	// Check if the returned ID is 0
	if id != 0 {
		t.Errorf("Expected ID 0 but got %d", id)
	}

	// Check if there's an error
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestWriteJSON(t *testing.T) {
	// Create a mock ResponseWriter
	w := httptest.NewRecorder()

	// Create mock data
	data := envelope{"key": "value"}

	// Call the helper function
	err := app.writeJSON(w, http.StatusOK, data, nil)

	// Check if there's no error
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	// Check if the response body is correct
	expectedBody := `{
		"key": "value"
	}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected response body %s but got %s", expectedBody, w.Body.String())
	}
}
