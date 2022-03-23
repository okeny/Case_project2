//This is a postgres repository using GORM

package repository

import (
	"fmt"
	"project-cases/entity"
)

type postRepo struct{}

//NewMemoryRepository
func NewPostgresRepository() CaseRepository {
	return &repo{}
}

//Save post to memory
func (*postRepo) Save(Case *entity.Case) (*entity.Case, error) {

	//Add the case to the collection
	return Case, nil
}

func (*postRepo) FindAll() ([]entity.Case, error) {
	return cases, nil
}

func (*postRepo) FindById(id int) (*entity.Case, error) {
	//loop through the cases to find one with the ID

	return &entity.Case{}, fmt.Errorf("Case with ID %v not found", id)
}

func (*postRepo) UpdateCase(c *entity.Case) (*entity.Case, error) {
	//loop throught the cases to find one to update
	
	return &entity.Case{}, fmt.Errorf("Case with ID %v not found", c.ID)
}

func (*postRepo) Delete(id int) error {


	return fmt.Errorf("Cases with ID %v not found", id)
}
