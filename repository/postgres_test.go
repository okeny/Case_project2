package repository

import (
	"project-cases/entity"
	mocks "project-cases/mocks/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

//Validation Test Cases
func TestPostSave(t *testing.T) {
	mockRepo := new(mocks.DatabaseInterface)
	Case := entity.Case{ID: 1, Title: "A", Description: "B"}

	//Expectations
	mockRepo.On("GetConnection").Return(func()(g *gorm.DB){ return g})
	mockRepo.On("Create",&Case).Return(nil)
	//creating testservice with a mock repo
	testRepo := NewPostgresRepository(mockRepo.GetConnection())
	result, err := testRepo.Save(&Case)

	//Data Assertions
	assert.Equal(t, Case.Title, result.Title)
	assert.Equal(t, Case.Description, result.Description)
	assert.Nil(t, err)
}

// //FindAll TestCases
func TestFindAll(t *testing.T) {

	//creating testservice with a mock repo
	testRepo := NewMemoryRepository()
	_, err := testRepo.FindAll()

	//Data Assertions
	assert.Nil(t, err)
}

//FindByID test Case
func TestFindByID(t *testing.T) {

	Case := entity.Case{ID: 1, Title: "A", Description: "B"}
	//creating testservice with a mock repo
	testRepo := NewMemoryRepository()
	result1, _ := testRepo.Save(&Case)
	result, err := testRepo.FindById(Case.ID)

	//Data Assertions
	assert.Nil(t, err)
	assert.Equal(t, result1.ID, result.ID)
	assert.Equal(t, result1.Title, result.Title)
	assert.Equal(t, result1.Description, result.Description)
}

func TestUpdate(t *testing.T) {

	Case := entity.Case{
		ID:          3,
		Title:       "came back",
		Description: "Bag stolen from Kira",
	}
	//creating testservice with a mock repo
	testRepo := NewMemoryRepository()
	result, err := testRepo.UpdateCase(&Case)

	//Data Assertions
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Case with ID 3 not found")
	assert.Equal(t, &entity.Case{}, result)
}

func TestDelete(t *testing.T) {

	Case := entity.Case{
		Title:       "came back",
		Description: "Bag stolen from Kira",
	}

	testRepo := NewMemoryRepository()
	res, _ := testRepo.Save(&Case)

	err := testRepo.Delete(res.ID)
	assert.Nil(t, err)
}
