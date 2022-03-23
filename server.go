package main

import (
	"log"
	"project-cases/controller"
	"project-cases/repository"
	"project-cases/service"

	"github.com/gin-gonic/gin"
)

var (
	CaseRepository repository.CaseRepository = repository.NewMemoryRepository()
	caseService service.CaseService = service.NewCaseService(CaseRepository)
	caseController controller.CaseController = controller.NewCaseCrontroller(caseService)
)

func main() {

	var address = ":3001"
	r := gin.Default()

	r.GET("/cases", caseController.GetCases)
	r.GET("/case/:id", caseController.GetCase)
	r.POST("/cases", caseController.PostCase)
	r.PUT("/cases", caseController.PutCase)
	r.DELETE("/cases/:id", caseController.DeleteCase)

	//_ = utility.GetConnection()
	log.Fatal(r.Run(address))
}
