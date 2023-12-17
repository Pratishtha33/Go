package datastore

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"xyz/model"
)

// leave is a struct representing the data store for leaves
type leave struct{}

// New creates a new instance of the leave data store
func New() *leave {
	return &leave{}
}

// GetByID retrieves a leave record by its ID.
func (s *leave) GetByID(ctx *gofr.Context, id string) (*model.Leave, error) {
	var resp model.Leave

	// Construct the SQL query to fetch a leave record by ID.
	query := "SELECT ID, EmployeeId, StartDate, EndDate, Reason FROM leaves WHERE ID=" + id

	// Execute the query and scan the result into the leave model.
	err := ctx.DB().QueryRowContext(ctx, query).
		Scan(&resp.ID, &resp.EmployeeID, &resp.StartDate, &resp.EndDate, &resp.Reason)

	switch err {
	case sql.ErrNoRows:
		return &model.Leave{}, errors.EntityNotFound{Entity: "leave", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Leave{}, err
	}
}

// Create adds a new leave record to the data store.
func (s *leave) Create(ctx *gofr.Context, leave *model.Leave) (*model.Leave, error) {
    var resp model.Leave

    // Insert a new leave record into the database.
    _, err := ctx.DB().ExecContext(ctx, "INSERT INTO leaves (EmployeeID, StartDate, EndDate, Reason) VALUES (?, ?, ?, ?);", leave.EmployeeID, leave.StartDate, leave.EndDate, leave.Reason)

    if err != nil {
        return &model.Leave{}, errors.DB{Err: err}
    }

    // Try to query the inserted record.
    err = ctx.DB().QueryRowContext(ctx, "SELECT ID, EmployeeID, StartDate, EndDate, Reason FROM leaves WHERE ID = LAST_INSERT_ID();").
        Scan(&resp.ID, &resp.EmployeeID, &resp.StartDate, &resp.EndDate, &resp.Reason)

    if err != nil {
        return &model.Leave{}, errors.DB{Err: err}
    }

    return &resp, nil
}



// Update modifies an existing leave record in the data stor
func (s *leave) Update(ctx *gofr.Context, leave *model.Leave) (*model.Leave, error) {
    // Construct the SQL query to update a leave record using parameterized query.
    query := "UPDATE leaves SET StartDate=?, EndDate=?, Reason=? WHERE ID=?"

    // Execute the update query with parameters
    _, err := ctx.DB().ExecContext(ctx, query, leave.StartDate, leave.EndDate, leave.Reason, leave.ID)
    if err != nil {
        return &model.Leave{}, errors.DB{Err: err}
    }

    return leave, nil
}



// Delete removes a leave record from the data store by its ID.
func (s *leave) Delete(ctx *gofr.Context, id int) error {
	// Construct the SQL query to delete a leave record by ID.
	query := "DELETE FROM leaves WHERE ID=" + strconv.Itoa(id)

	// Execute the delete query.
	_, err := ctx.DB().ExecContext(ctx, query)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
