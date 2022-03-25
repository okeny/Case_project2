package repository

import (
	"project-cases/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Validation Test Cases
func TestSave(t *testing.T) {

	Case := entity.Case{ID: 1, Title: "A", Description: "B"}

	//creating testRepo
	testRepo := NewMemoryRepository()
	result, err := testRepo.Save(&Case)

	//Data Assertions
	assert.Equal(t, Case.Title, result.Title)
	assert.Equal(t, Case.Description, result.Description)
	assert.Nil(t, err)
}

// //FindAll TestCases
func TestFindAllWithEmptyMemory(t *testing.T) {

	//creating testservice with a mock repo
	testRepo := NewMemoryRepository()
	_, err := testRepo.FindAll()

	//Data Assertions
	assert.Nil(t, err)
}

// //FindAll TestCases
func TestFindAllWithData(t *testing.T) {

	Case := entity.Case{Title: "A", Description: "B"}

	//creating testRepo
	testRepo := NewMemoryRepository()
	result1, _ := testRepo.Save(&Case)
	result2, err := testRepo.FindAll()

	//Data Assertions
	assert.Nil(t, err)
	assert.Equal(t, result1.Title, result2[0].Title)
	assert.Equal(t, result1.Description, result2[0].Description)
}

//FindByID test Case
func TestFindByIDInTheMemory(t *testing.T) {

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

func TestFindByIDNotInTheMemory(t *testing.T) {

	//creating testservice with a mock repo
	testRepo := NewMemoryRepository()
	_, err := testRepo.FindById(1)

	//Data Assertions
	assert.NotNil(t, err)
	assert.Equal(t, "Case with ID 1 not found", err.Error())
}

func TestUpdateWithIDNotInMemory(t *testing.T) {

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

func TestUpdateWithIDInMemory(t *testing.T) {

	Case := entity.Case{
		Title:       "came back",
		Description: "Bag stolen from Kira",
	}

	//creating testservice with a mock repo
	testRepo := NewMemoryRepository()
	res, _ := testRepo.Save(&Case)

	Case2 := entity.Case{
		ID:          res.ID,
		Title:       "came back Soon",
		Description: "Bag stolen from Kampala",
	}

	result, err := testRepo.UpdateCase(&Case2)

	//Data Assertions
	assert.Nil(t, err)
	assert.Equal(t, Case2.ID, result.ID)
	assert.Equal(t, Case2.Title, result.Title)
	assert.Equal(t, Case2.Description, result.Description)
}
func TestDeleteWithReturnedID(t *testing.T) {

	Case := entity.Case{
		Title:       "came back",
		Description: "Bag stolen from Kira",
	}

	testRepo := NewMemoryRepository()
	res, _ := testRepo.Save(&Case)

	err := testRepo.Delete(res.ID)
	assert.Nil(t, err)
}

func TestDeleteWithIdNotInTheMemory(t *testing.T) {

	testRepo := NewMemoryRepository()

	err := testRepo.Delete(20)
	assert.NotNil(t, err)
	assert.Equal(t, "Cases with ID 20 not found", err.Error())
}
