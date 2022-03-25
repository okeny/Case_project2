//This is a postgres repository using GORM

package repository

import (
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

	err := p.db.Create(&Case).Error
	if err != nil {
		return nil, err
	}
	//Add the case to the collection
	return Case, nil
}

func (p *postRepo) FindAll() ([]entity.Case, error) {
	var Cases []entity.Case
	err := p.db.Select("id, title, description").Order("created_at desc").Find(&Cases).Error
	if err != nil {
		return nil, err
	}
	return Cases, nil
}

func (p *postRepo) FindById(id int) (*entity.Case, error) {
	//loop through the cases to find one with the ID
	var Case *entity.Case
	err := p.db.Select("id, title, description").First(&Case, id).Error
	if err != nil {
		return nil, err
	}
	return Case, nil
}

func (p *postRepo) UpdateCase(c *entity.Case) (*entity.Case, error) {
	//loop throught the cases to find one to update
	err := p.db.Save(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (p *postRepo) Delete(id int) error {

	err := p.db.Delete(&entity.Case{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
