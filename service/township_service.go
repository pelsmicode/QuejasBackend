package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type TownshipService interface {
	GetTownships() ([]model.Township, error)
	GetTownshipByID(id int) (model.Township, error)
}

type DefaultTownshipServcie struct {
	R repository.TownshipeRepository
}

func NewTownshipService(r repository.TownshipeRepository) DefaultTownshipServcie {
	return DefaultTownshipServcie{R: r}
}

func (s DefaultTownshipServcie) GetTownships() ([]model.Township, error) {
	return s.R.FindAll()
}

func (s DefaultTownshipServcie) GetTownshipByID(id int) (model.Township, error) {
	return s.R.FindByID(id)
}
