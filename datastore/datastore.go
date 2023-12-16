package datastore

import (
	"strconv"
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"simple-rest-api/model"
)

type student struct{}

func New() *student {
	return &student{}
}

func (s *student) GetByID(ctx *gofr.Context, id string) (*model.Student, error) {
	var resp model.Student

	query := "SELECT id, name, age, class FROM students WHERE id=" + id

	err := ctx.DB().QueryRowContext(ctx, query).
		Scan(&resp.ID, &resp.Name, &resp.Age, &resp.Class)

	switch err {
	case sql.ErrNoRows:
		return &model.Student{}, errors.EntityNotFound{Entity: "student", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Student{}, err
	}
}

func (s *student) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	var resp model.Student
 
	err := ctx.DB().QueryRowContext(ctx, "INSERT INTO students (name, age, class) VALUES (?, ?, ?);", student.Name, student.Age, student.Class).Scan(&resp.ID, &resp.Name, &resp.Age, &resp.Class)
 
	if err != nil {
	   return &model.Student{}, errors.DB{Err: err}
	}
 
	return &resp, nil
 }

 func (s *student) Update(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
    query := "UPDATE students SET name='" + student.Name + "', age=" + strconv.Itoa(student.Age) +
        ", class='" + student.Class + "' WHERE id=" + strconv.Itoa(student.ID)

    _, err := ctx.DB().ExecContext(ctx, query)
    if err != nil {
        return &model.Student{}, errors.DB{Err: err}
    }

    return student, nil
}



func (s *student) Delete(ctx *gofr.Context, id int) error {
	query := "DELETE FROM students WHERE id=" + strconv.Itoa(id)

	_, err := ctx.DB().ExecContext(ctx, query)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}