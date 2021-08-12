package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type DepartmentRepository struct {
	client *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) DepartmentRepository {
	return DepartmentRepository{db}
}

func (r *DepartmentRepository) FindAll() ([]model.Department, error) {
	query := `SELECT id,name FROM departments`

	var departments []model.Department
	err := r.client.Select(&departments, query)
	if err != nil {
		log.Println("DepartmentRepository\t [DB Department Error]", err)
		return nil, err
	}

	return departments, nil
}

func (r *DepartmentRepository) FindByID(id int) (model.Department, error) {
	query := `SELECT id,name FROM departments WHERE id=$1`

	var department model.Department
	err := r.client.Get(&department, query, id)
	if err != nil {
		log.Println("DepartmentRepository\t [DB Department Error]", err)
		return department, err
	}

	return department, nil
}
