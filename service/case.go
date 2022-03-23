package service

import (
	"errors"
	"math/rand"
	"project-cases/entity"
	"project-cases/repository"
)

type CaseService interface {
	Validate(Case *entity.Case) error
	Create(Case *entity.Case) (*entity.Case, error)
	FindAll() ([]entity.Case, error)
	FindById(id int) (*entity.Case, error)
	Update(Case *entity.Case) (*entity.Case, error)
	Delete(id int) error
}

type service struct{ 
	repo repository.CaseRepository
 }

func NewCaseService(rep repository.CaseRepository) CaseService {
	return &service{
		repo: rep,
	}
}

func (*service) Validate(Case *entity.Case) error {
	//validate the id
	if Case == nil {
		return errors.New("The Case Params are empty")
	}

	if Case.ID != 0 {
		return errors.New("case must not include an ID")
	}

	if Case.Title == "" {
		return errors.New("Case must include a title")
	}
	return nil
}

func (s *service) Create(Case *entity.Case) (*entity.Case, error) {
	Case.ID = int(rand.Int63())
	return s.repo.Save(Case)
}

func (s *service) FindAll() ([]entity.Case, error) {
	return s.repo.FindAll()
}

func (s *service) FindById(id int) (*entity.Case, error) {
	return s.repo.FindById(id)
}

func (s *service) Update(Case *entity.Case) (*entity.Case, error) {
	return s.repo.UpdateCase(Case)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
