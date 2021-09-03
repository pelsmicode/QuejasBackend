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

func (r *ComplaintRepository) GetMainComplaints() ([]model.Complaint, error) {
	query := `select 
	complaints.id, complaints.no_doc, complaints.date_doc, complaints.created_at, companies.id "company.id", companies.name "company.name",
    companies.township_id "company.township.id", companies.name "company.name", persons.id "person.id", persons.name "person.name", persons.lastname "person.lastname", diaco_branches.id "person.branch.id", 
    diaco_branches.name "person.branch.name", townships.name "person.branch.township.name", diaco_branches.name "person.branch.township.department.name"
from complaints
inner join companies 
on complaints.company_id = companies.id
inner join persons 
on complaints.person_id = persons.id
inner join diaco_branches
on persons.branch_id = diaco_branches.id
inner join townships 
on diaco_branches.township_id = townships.id
inner join departments 
on townships.department_id = departments.id`

	var c []model.Complaint
	err := r.client.Select(&c, query)
	if err != nil {
		log.Println("ComplaintRepository\t [DB Complaint Error]", err)
		return nil, err
	}

	return c, nil
}
