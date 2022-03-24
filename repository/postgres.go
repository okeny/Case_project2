//This is a postgres repository using GORM

package repository

import (
	"fmt"
	"project-cases/entity"

	"gorm.io/gorm"
)

type postRepo struct{ db *gorm.DB }

//NewMemoryRepository
func NewPostgresRepository(db *gorm.DB) CaseRepository {
	return &postRepo{db}
}

//Save post to memory
func (p *postRepo) Save(Case *entity.Case) (*entity.Case, error) {

	result := p.db.Create(&Case)
	fmt.Println(&result.Statement)
	//Add the case to the collection
	return &entity.Case{}, nil
}

func (p *postRepo) FindAll() ([]entity.Case, error) {
	var Case *entity.Case
	result := p.db.Select("id, title, description").Order("created_at desc").Find(&Case)
	fmt.Println(result)
	return cases, nil
}

func (p *postRepo) FindById(id int) (*entity.Case, error) {
	//loop through the cases to find one with the ID
	var Case *entity.Case
	result := p.db.Select("id, title, description").First(&Case, id)
	fmt.Println(result)
	return &entity.Case{}, fmt.Errorf("Case with ID %v not found", id)
}

func (p *postRepo) UpdateCase(c *entity.Case) (*entity.Case, error) {
	//loop throught the cases to find one to update
	p.db.Save(&c)
	return &entity.Case{}, fmt.Errorf("Case with ID %v not found", c.ID)
}

func (p *postRepo) Delete(id int) error {

	p.db.Delete(&entity.Case{}, id)
	return fmt.Errorf("Cases with ID %v not found", id)
}
