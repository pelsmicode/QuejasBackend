package service

import (
	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/repository"
)

type BranchService interface {
	GetBranches() ([]model.DiacoBranch, error)
}

type DefaultBranchService struct {
	R repository.BranchRepository
}

func NewBranchService(r repository.BranchRepository) DefaultBranchService {
	return DefaultBranchService{R: r}
}

func (s DefaultBranchService) GetBranches() ([]model.DiacoBranch, error) {
	return s.R.FindAll()
}
