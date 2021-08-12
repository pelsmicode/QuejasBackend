package repository

import "github.com/jmoiron/sqlx"

type DepartmentRepository struct {
	client *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) DepartmentRepository {
	return DepartmentRepository{db}
}
