package datastore

import (
	"xyz/model"

	"gofr.dev/pkg/gofr"
)

type Leave interface {
	// GetByID retrieves a leave record based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.Leave, error)
	//  inserts a new leave record into the database.
	Create(ctx *gofr.Context, leave *model.Leave) (*model.Leave, error)
	// updates an existing leave record with the provided information.
	Update(ctx *gofr.Context, leave *model.Leave) (*model.Leave, error)
	// Delete  a leave record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}
