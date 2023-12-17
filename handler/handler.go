package leavemanagementhandler

import (
	"strconv" // // Package strconv implements conversions to and from string representations
	// of basic data types

	"gofr.dev/pkg/errors" // Package errors implements functions to manipulate errors and wrap them
	// with additional information.
	"gofr.dev/pkg/gofr"  // Package gofr is a custom package that likely provides utility functions

	"simple-rest-api/datastore"  // The "simple-rest-api/datastore" package contains implementations of data storage and retrieval for the application
	"simple-rest-api/model"
)

type handler struct {
	store datastore.Leave
}

// New initializes a new handler with the provided leave data store.
func New(s datastore.Leave) handler {
	return handler{store: s}
}

// GetByID retrieves a leave record by its ID.
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	// Extract leave ID from the request path parameters.
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	// Validate that the ID is a valid integer.
	if _, err := validateID(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	// Fetch leave record from the data store.
	resp, err := h.store.GetByID(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "leave",
			ID:     id,
		}
	}

	return resp, nil
}

// Create creates a new leave record based on the provided data.
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var leave model.Leave

	// Bind the incoming JSON data to the leave model.
	if err := ctx.Bind(&leave); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	// Insert the new leave record into the data store.
	resp, err := h.store.Create(ctx, &leave)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Update updates an existing leave record with the provided data.
func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	// Extract leave ID from the request path parameters.
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	// Validate that the ID is a valid integer.
	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var leave model.Leave
	// Bind the incoming JSON data to the leave model.
	if err = ctx.Bind(&leave); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	// Set the ID for the leave record.
	leave.ID = id

	// Update the existing leave record in the data store.
	resp, err := h.store.Update(ctx, &leave)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Delete removes a leave record from the data store based on its ID.
func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	// Extract leave ID from the request path parameters.
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	// Validate that the ID is a valid integer.
	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	// Delete the leave record from the data store.
	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}

	return "Deleted successfully", nil
}

// validateID converts a string ID to an integer and validates it.
func validateID(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return res, nil
}
