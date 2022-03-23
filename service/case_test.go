package service

import (
	"project-cases/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)


//Validation Test Cases
func TestValidateEmptyCase(t *testing.T) {

	testService := NewCaseService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The Case Params are empty", err.Error())

}

func TestValidateIncludedID(t *testing.T) {
	Case := entity.Case{
		ID:          3,
		Title:       "Stolen Bag",
		Description: "Bag stolen from Kira",
	}

	testService := NewCaseService(nil)
	err := testService.Validate(&Case)
	assert.NotNil(t, err)
	assert.Equal(t, "case must not include an ID", err.Error())

}

func TestValidateEmptyTitle(t *testing.T) {

	Case := entity.Case{
		Title:       "",
		Description: "Bag stolen from Kira",
	}

	testService := NewCaseService(nil)
	err := testService.Validate(&Case)
	assert.NotNil(t, err)
	assert.Equal(t, "Case must include a title", err.Error())
}
//still debuging
//Create TestCase

// func TestCreate(t *testing.T) {

// 	mockRepo := new(MockRepository)

// 	Case := entity.Case{ID: 1, Title: "A", Description: "B"}
// 	//Expectations
// 	mockRepo.On("Save").Return(&Case, nil)
// 	//creating testservice with a mock repo
// 	testService := NewCaseService(mockRepo)
// 	result, err := testService.Create(&Case)

// 	mockRepo.AssertExpectations(t)

// 	//Data Assertions
// 	assert.Equal(t, Case.Title, result.Title)
// 	assert.Equal(t, Case.Description, result.Description)
// 	assert.Nil(t, err)
// }

// //FindAll TestCases
// func TestFindAll(t *testing.T) {
// 	mockRepo := new(MockRepository)

// 	Case := entity.Case{ID: 1, Title: "A", Description: "B"}
// 	//Expectations
// 	mockRepo.On("FindAll").Return([]entity.Case{Case}, nil)
// 	//creating testservice with a mock repo
// 	testService := NewCaseService(mockRepo)
// 	result, err := testService.FindAll()

// 	mockRepo.AssertExpectations(t)

// 	//Data Assertions
// 	assert.Equal(t, Case.Title, result[0].Title)
// 	assert.Equal(t, Case.Description, result[0].Description)
// 	assert.Nil(t, err)

// }

// //FindByID test Case
// func TestFindByID(t *testing.T){
	
// 	mockRepo := new(MockRepository)

// 	Case := entity.Case{ID: 1, Title: "A", Description: "B"}
// 	//Expectations
// 	mockRepo.On("FindByID").Return(&Case, nil)
// 	//creating testservice with a mock repo
// 	testService := NewCaseService(mockRepo)
// 	result, _ := testService.FindById(Case.ID)
	
// 	mockRepo.AssertExpectations(t)

// 	//Data Assertions
// 	assert.Equal(t, Case.ID, result.ID)
// 	assert.Equal(t, Case.Title, result.Title)
// 	assert.Equal(t, Case.Description, result.Description)
// }

// func TestUpdate(t *testing.T){
	
// 	mockRepo := new(MockRepository)

// 	Case := entity.Case{ID: 1, Title: "A", Description: "B"}
// 	//Expectations
// 	mockRepo.On("UpdateCase").Return(&Case, nil)
// 	//creating testservice with a mock repo
// 	testService := NewCaseService(mockRepo)
// 	result, _ := testService.Create(&Case)
// 	mockRepo.AssertExpectations(t)

// 	//Data Assertions
// 	assert.Equal(t, Case.ID, result.ID)
// 	assert.Equal(t, Case.Title, result.Title)
// 	assert.Equal(t, Case.Description, result.Description)
// }

// func TestDelete(t *testing.T){
	
// 	mockRepo := new(MockRepository)

// 	//Expectations
// 	mockRepo.On("Delete").Return(nil)
// 	//creating testservice with a mock repo
// 	testService := NewCaseService(mockRepo)
// 	err := testService.Delete(1)
// 	mockRepo.AssertExpectations(t)

// 	//Data Assertions
// 	assert.Nil(t, err)
// }



