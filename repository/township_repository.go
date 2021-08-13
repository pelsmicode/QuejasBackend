package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type TownshipeRepository struct {
	client *sqlx.DB
}

func NewTownshipRepository(db *sqlx.DB) TownshipeRepository {
	return TownshipeRepository{db}
}

func (r *TownshipeRepository) FindAll() ([]model.Township, error) {
	query := `SELECT id,name FROM townships`

	var townships []model.Township
	err := r.client.Select(&townships, query)
	if err != nil {
		log.Println("TownshipeRepository\t [DB Township Error]", err)
		return nil, err
	}

	return townships, nil
}

func (r *TownshipeRepository) FindByID(id int) (model.Township, error) {
	query := `SELECT townships.id, townships.name,
						departments.id "department.id", departments.name "department.name",
						regions.id "department.region.id", regions.name "department.region.name" 
						FROM townships 
						INNER JOIN departments
						on townships.department_id = departments.id
						INNER JOIN regions
						on regions.id = departments.region_id
						WHERE townships.id=$1`

	var muni model.Township
	err := r.client.Get(&muni, query, id)
	if err != nil {
		log.Println("TownshipeRepository\t [DB Township Error]", err)
		return muni, err
	}

	return muni, nil
}

func (r *TownshipeRepository) FindByDeparment(depId int) ([]model.Township, error) {
	query := `SELECT townships.id, townships.name,
						departments.id "department.id", departments.name "department.name",
						regions.id "department.region.id", regions.name "department.region.name"
						FROM townships
						LEFT JOIN departments
						ON townships.department_id = departments.id
						INNER JOIN regions
						ON regions.id = departments.region_id
						WHERE departments.id=$1`

	var townships []model.Township
	err := r.client.Select(&townships, query, depId)
	if err != nil {
		log.Println("TownshipeRepository\t [DB Township Error]", err)
		return nil, err
	}

	return townships, nil
}
