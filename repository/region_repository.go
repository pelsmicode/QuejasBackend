package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type RegionRepository struct {
	client *sqlx.DB
}

func NewRegionRepository(db *sqlx.DB) RegionRepository {
	return RegionRepository{db}
}

func (r *RegionRepository) FindAll() ([]model.Region, error) {
	query := `SELECT id,name FROM regions`

	var regions []model.Region
	err := r.client.Select(&regions, query)
	if err != nil {
		log.Println("RegionRepository\t [DB Region Error]", err)
		return nil, err
	}

	return regions, nil
}

func (r *RegionRepository) FindByID(id int) (model.Region, error) {
	query := `SELECT id,name FROM regions WHERE id=$1`

	var region model.Region
	err := r.client.Get(&region, query, id)
	if err != nil {
		log.Println("RegionRepository\t [DB Region Error]", err)
		return region, nil
	}

	return region, nil
}
