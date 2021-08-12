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
	r repository.RegionRepository
}

func NewRegionService(r repository.RegionRepository) DefaultRegionService {
	return DefaultRegionService{r: r}
}

func (s *DefaultRegionService) GetRegions() ([]model.Region, error) {
	return s.r.FindAll()
}

func (s *DefaultRegionService) GetRegionByID(id int) (model.Region, error) {
	return s.r.FindByID(id)
}
