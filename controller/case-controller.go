package controller

import (
	"net/http"
	"project-cases/entity"
	"project-cases/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CaseController interface{
	GetCases(c *gin.Context)
	DeleteCase(c *gin.Context)
	GetCase(c *gin.Context)
	PostCase(c *gin.Context)
	PutCase(c *gin.Context)
}

type controller struct{}

var (
	caseService service.CaseService 
)

func NewCaseCrontroller(service service.CaseService) CaseController{
	caseService = service
	return &controller{}
}
func (*controller) GetCases(c *gin.Context) {

	cases, err := caseService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cases)
}

// DELETE /books/:id
// Deleting a case
func (*controller) DeleteCase(c *gin.Context) {

	//Validating input
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err1 := caseService.Delete(id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, "item deleted")
}

func (*controller) GetCase(c *gin.Context) {

	id := c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cases, err2 := caseService.FindById(id2)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(http.StatusOK, cases)
}

func (*controller) PostCase(c *gin.Context) {

	var payload entity.Case
	if e := c.ShouldBindJSON(&payload); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}
	//validate the id
	if payload.ID != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "New case must not includean ID"})
		return
	}

	cases, err := caseService.Create(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Case Added", "data": cases})
}

func (*controller) PutCase(c *gin.Context) {

	var payload entity.Case
	if e := c.ShouldBindJSON(&payload); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	_, err := caseService.Update(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Case Updated Successfully", "data": ""})
}
