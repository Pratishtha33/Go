package main

import (
	"xyz/datastore"         // Import the leave datastore package.
	"xyz/handler" // Import the leave management handler package.

	"gofr.dev/pkg/gofr"
)

func main() {
	// Initialize a new GoFr application.
	app := gofr.New()

	// Create a new instance of the employee leave datastore.
	s := datastore.New()

	// Create a new instance of the employee leave management handler.
	h := handler.New(s)

	
	// different HTTP methods to perform CRUD operations
	app.GET("/leaves/{id}", h.GetByID)   // Retrieve leave details by ID
	app.POST("/leaves", h.Create)        // Create a new leave record
	app.PUT("/leaves/{id}", h.Update)    // Update an existing leave record
	app.DELETE("/leaves/{id}", h.Delete) // Delete a leave record

	// Start the server on a custom port
	app.Server.HTTP.Port = 8000
	app.Start()
}
