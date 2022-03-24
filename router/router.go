package router

import (
	"project-cases/controller"
	//"project-cases/database"
	"project-cases/repository"
	"project-cases/service"

	"github.com/gin-gonic/gin"
)

var (

	//Postgres Repo
	//db  = database.NewDatabase("localhost", 5432, "postgres", "root", "caseproject")
	//CaseRepository repository.CaseRepository = repository.NewPostgresRepository(db)
	//Memory Repo
	CaseRepository repository.CaseRepository = repository.NewMemoryRepository()
	//Other Params(setting up controller and service)
	caseService    service.CaseService       = service.NewCaseService(CaseRepository)
	caseController controller.CaseController = controller.NewCaseCrontroller(caseService)
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/cases", caseController.GetCases)
	r.GET("/case/:id", caseController.GetCase)
	r.POST("/cases", caseController.PostCase)
	r.PUT("/cases", caseController.PutCase)
	r.DELETE("/cases/:id", caseController.DeleteCase)

	return r
}
