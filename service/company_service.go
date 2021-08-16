package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type CompanyService interface {
	SaveCompany(model.CompanyRequest) (int, error)
	UpdateCompany(model.CompanyRequest) (model.ComapnyResponse, error)
}

type DefaultCompanyService struct {
	R repository.CompanyRepository
}

func NewCompanyService(r repository.CompanyRepository) DefaultCompanyService {
	return DefaultCompanyService{R: r}
}

func (s DefaultCompanyService) SaveCompany(c model.CompanyRequest) (int, error) {
	return s.R.Save(c)
}

func (s DefaultCompanyService) UpdateCompany(c model.CompanyRequest) (model.ComapnyResponse, error) {
	return s.R.Update(c)
}
