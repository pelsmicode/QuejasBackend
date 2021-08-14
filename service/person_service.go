package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type PersonService interface {
	SavePerson(model.PersonRequest) error
}

type DefaultPersonService struct {
	R repository.PersonRepository
}

func NewPersonService(r repository.PersonRepository) DefaultPersonService {
	return DefaultPersonService{R: r}
}

func (s DefaultPersonService) SavePerson(p model.PersonRequest) error {
	return s.R.Save(p)
}
