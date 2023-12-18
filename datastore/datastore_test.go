package leavedatastore

import (
	"database/sql"
	"errors"
	"testing"

	"gofr.dev/pkg/gofr"

	"gofr.dev/model"
)

// MockDB is a mock implementation of the database interface for testing.
type MockDB struct{}

func (db *MockDB) QueryRowContext(ctx *gofr.Context, query string, args ...interface{}) *sql.Row {
	
	return nil
}

func (db *MockDB) ExecContext(ctx *gofr.Context, query string, args ...interface{}) (sql.Result, error) {
	
	return nil, nil
}

func TestGetByID(t *testing.T) {
	// Create a new instance of the leave data store with a mock database.
	store := &leave{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a leave ID for testing.
	id := "1"

	// Call the GetByID method.
	_, err := store.GetByID(ctx, id)

	// Check if the error is the expected "EntityNotFound" error.
	if !errors.Is(err, errors.New("EntityNotFound")) {
		t.Errorf("Expected EntityNotFound error, got %v", err)
	}
}

func TestCreate(t *testing.T) {
	// Create a new instance of the leave data store with a mock database.
	store := &leave{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a leave record for testing.
	newLeave := &model.Leave{EmployeeID: 123, StartDate: "2023-01-01", EndDate: "2023-01-05", Reason: "Vacation"}

	// Call the Create method.
	_, err := store.Create(ctx, newLeave)

	// Check if there is no error.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	// Create a new instance of the leave data store with a mock database.
	store := &leave{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a leave record for testing.
	updatedLeave := &model.Leave{ID: 1, StartDate: "2023-02-01", EndDate: "2023-02-05", Reason: "Updated Vacation"}

	// Call the Update method.
	_, err := store.Update(ctx, updatedLeave)

	// Check if there is no error.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDelete(t *testing.T) {
	// Create a new instance of the leave data store with a mock database.
	store := &leave{db: &MockDB{}}

	// Mock a context with a mock database.
	ctx := &gofr.Context{DBProvider: func() gofr.DB { return &MockDB{} }}

	// Mock a leave ID for testing.
	leaveID := 1

	// Call the Delete method.
	err := store.Delete(ctx, leaveID)

	// Check if there is no error.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
