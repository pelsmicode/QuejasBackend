package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type CompanyRepository struct {
	client *sqlx.DB
}

func NewCompanyRepository(db *sqlx.DB) CompanyRepository {
	return CompanyRepository{db}
}

func (r *CompanyRepository) Save(company model.CompanyRequest) error {
	query := `INSERT INTO companies (name, township_id, nit, address, phone, email)
	VALUES ($1, $2, $3, $4, $5, $6)`

	tx := r.client.MustBegin()
	tx.MustExec(query, company.Name, company.Township, company.NIT, company.Addres, company.Phone, company.Email)
	if err := tx.Commit(); err != nil {
		log.Println("CompanyRepository\t [DB Companies Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}

func (r *CompanyRepository) Update(company model.CompanyRequest) (model.ComapnyResponse, error) {
	query := `UPDATE companies
	SET name=$1,
	township_id=$2,
	nit=$3,
	address=$4,
	phone=$5,
	email=$6
	WHERE id=$7`

	tx := r.client.MustBegin()
	tx.MustExec(query, company.Name, company.Township, company.NIT, company.Addres, company.Phone, company.Email, company.ID)
	if err := tx.Commit(); err != nil {
		log.Println("CompanyRepository\t [DB Companies Error]", err)
		tx.Rollback()
		return model.ComapnyResponse{}, err
	}

	return model.ToCompanyReponse(company), nil
}

func (r *CompanyRepository) GetLastID() int {
	var id int
	err := r.client.Get(&id, "select max(id) from companies")
	if err != nil {
		log.Println("CompanyRepository\t [DB Companies Error]", err)
		return 0
	}

	return id
}
