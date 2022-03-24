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

type controller struct{ 
	service service.CaseService
}

//This setup the service
func NewCaseCrontroller(service service.CaseService) CaseController{
	return &controller{
		service,
	}
}

// GET /cases
// Get a cases
//Returns a status code or an error
func (ct *controller) GetCases(c *gin.Context) {

	cases, err := ct.service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cases)
}

// DELETE /case/:id
// Deleting a case
func (ct *controller) DeleteCase(c *gin.Context) {

	//Validating input
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err1 := ct.service.Delete(id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, "item deleted")
}

// GET /case/:id
// Getting a single case
func (ct *controller) GetCase(c *gin.Context) {

	id := c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cases, err2 := ct.service.FindById(id2)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(http.StatusOK, cases)
}
//POST /case
//
func (ct *controller) PostCase(c *gin.Context) {

	var payload entity.Case
	if e := c.ShouldBindJSON(&payload); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}
	
	cases, err := ct.service.Create(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Case Added", "data": cases})
}
//Updating a case
// endPoint /cases
func (ct *controller) PutCase(c *gin.Context) {

	var payload entity.Case
	if e := c.ShouldBindJSON(&payload); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	_, err := ct.service.Update(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Case Updated Successfully", "data": ""})
}
