package repository

import (
	"fmt"
	"project-cases/entity"
)

type repo struct{}

//NewMemoryRepository
func NewMemoryRepository() CaseRepository {
	return &repo{}
}

//Save post to memory
func (*repo) Save(Case *entity.Case) (*entity.Case, error) {

	//Add the case to the collection
	cases = append(cases, *Case)
	return Case, nil
}

func (*repo) FindAll() ([]entity.Case, error) {
	return cases, nil
}

func (*repo) FindById(id int) (*entity.Case, error) {
	//loop through the cases to find one with the ID
	for _, c := range cases {
		if c.ID == id {
			return &c, nil
		}
	}
	return &entity.Case{}, fmt.Errorf("Case with ID %v not found", id)
}

func (*repo) UpdateCase(c *entity.Case) (*entity.Case, error) {
	//loop throught the cases to find one to update
	for i, case1 := range cases {
		if case1.ID == c.ID {
			cases[i] = *c
			return c, nil
		}
	}
	return &entity.Case{}, fmt.Errorf("Case with ID %v not found", c.ID)
}

func (*repo) Delete(id int) error {
	//loop through the cases to find one to delete
	for i, u := range cases {
		if u.ID == id {
			//copy the case slice from the begining to the index and
			//plus one after
			cases = append(cases[:i], cases[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Cases with ID %v not found", id)
}
