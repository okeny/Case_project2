package service

import (
	"math/rand"
	"project-cases/entity"
	"project-cases/repository"
)

type CaseService interface {
	Create(Case *entity.Case) (*entity.Case, error)
	FindAll() ([]entity.Case, error)
	FindById(id int) (*entity.Case, error)
	Update(Case *entity.Case) (*entity.Case, error)
	Delete(id int) error
}

type service struct{}

var (
	repo repository.CaseRepository
)

func NewCaseService(rep repository.CaseRepository) CaseService {
	repo =rep
	return &service{}
}

func (*service) Create(Case *entity.Case) (*entity.Case, error) {
	Case.ID = int(rand.Int63())
	return repo.Save(Case)
}

func (*service) FindAll() ([]entity.Case, error) {
	return repo.FindAll()
}

func (*service) FindById(id int) (*entity.Case, error) {
    return repo.FindById(id)
}

func (*service) Update(Case *entity.Case) (*entity.Case, error) {
    return repo.UpdateCase(Case)
}

func (*service) Delete(id int) error {
   return repo.Delete(id)
}
