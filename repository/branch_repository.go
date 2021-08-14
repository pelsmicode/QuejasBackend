package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type BranchRepository struct {
	client *sqlx.DB
}

func NewBranchRepository(db *sqlx.DB) BranchRepository {
	return BranchRepository{db}
}

func (r *BranchRepository) FindAll() ([]model.DiacoBranch, error) {
	query := `SELECT diaco_branches.id, diaco_branches.name,
	townships.id "township.id", townships.name "township.name"
	FROM diaco_branches
	INNER JOIN townships 
	ON diaco_branches.township_id = townships.id`

	var branches []model.DiacoBranch
	err := r.client.Select(&branches, query)
	if err != nil {
		log.Println("BranchRepository\t [DB DiacoBranch Error]", err)
		return nil, err
	}

	return branches, nil
}
