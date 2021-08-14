package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type PersonRepository struct {
	client *sqlx.DB
}

func NewPersonRepository(db *sqlx.DB) PersonRepository {
	return PersonRepository{db}
}

func (r *PersonRepository) Save(person model.PersonRequest) error {
	query := `INSERT INTO persons (name, lastname, dpi, nit, gender, email, phone, township_id, branch_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	tx := r.client.MustBegin()
	tx.MustExec(query, person.Name, person.Lastname, person.DPI, person.NIT, person.Gender, person.Email, person.Phone, person.Township, person.Branch)
	if err := tx.Commit(); err != nil {
		log.Println("PersonRepository\t [DB Persons Error]", err)
		return err
	}

	return nil
}
