package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type ComplaintService interface {
	SaveComplaint(model.ComplaintRequest) error
	GetMainComplaints() ([]model.Complaint, error)
}

type DefaultComplaintService struct {
	R repository.ComplaintRepository
}

func NewComplaintService(r repository.ComplaintRepository) DefaultComplaintService {
	return DefaultComplaintService{R: r}
}

func (s DefaultComplaintService) SaveComplaint(c model.ComplaintRequest) error {
	return s.R.Save(c)
}

func (s DefaultComplaintService) GetMainComplaints() ([]model.Complaint, error) {
	return s.R.GetMainComplaints()
}
