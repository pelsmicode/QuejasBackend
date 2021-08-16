package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pelsmicode/QuejasBackend/model"
)

type ComplaintRepository struct {
	client *sqlx.DB
}

func NewComplaintRepository(db *sqlx.DB) ComplaintRepository {
	return ComplaintRepository{db}
}

func (r *ComplaintRepository) Save(complaint model.ComplaintRequest) error {
	query := `INSERT INTO complaints 
(no_doc, date_doc, detail, petition, company_id, complaints.person_id)
VALUES ($1,$2,$3,$4,$5,$6)`

	tx := r.client.MustBegin()
	tx.MustExec(query, complaint.NoDoc, complaint.DateDoc, complaint.Detail, complaint.Petition, complaint.Company, complaint.Person)
	if err := tx.Commit(); err != nil {
		log.Println("ComplaintRepository\t [DB Complaint Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}
