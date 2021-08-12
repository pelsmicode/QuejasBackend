package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type RegionService interface {
	GetRegions() ([]model.Region, error)
	GetRegionByID(id int) (model.Region, error)
}

type DefaultRegionService struct {
	R repository.RegionRepository
}

func NewRegionService(r repository.RegionRepository) DefaultRegionService {
	return DefaultRegionService{R: r}
}

func (s DefaultRegionService) GetRegions() ([]model.Region, error) {
	return s.R.FindAll()
}

func (s DefaultRegionService) GetRegionByID(id int) (model.Region, error) {
	return s.R.FindByID(id)
}
