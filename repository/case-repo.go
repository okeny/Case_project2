package repository

import "project-cases/entity"

var (
	cases []entity.Case
)

type CaseRepository interface {
	Save(Case *entity.Case) (*entity.Case, error)
	FindAll() ([]entity.Case, error)
	FindById(id int) (*entity.Case, error)
	UpdateCase(c *entity.Case) (*entity.Case, error)
	Delete(id int) error
}
