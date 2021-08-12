package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type DepartmentService interface {
	GetDeparments() ([]model.Department, error)
	GetDeparmentByID(id int) (model.Department, error)
}

type DefaultDepartmentService struct {
	R repository.DepartmentRepository
}

func NewDeparmentService(r repository.DepartmentRepository) DefaultDepartmentService {
	return DefaultDepartmentService{R: r}
}

func (s DefaultDepartmentService) GetDeparments() ([]model.Department, error) {
	return s.R.FindAll()
}

func (s DefaultDepartmentService) GetDeparmentByID(id int) (model.Department, error) {
	return s.R.FindByID(id)
}
