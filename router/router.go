package router

import (
	"os"
	"project-cases/config"
	"project-cases/controller"
	"project-cases/database"
	"project-cases/repository"
	"project-cases/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//Load Config
	config.Init()
     port,_ := strconv.Atoi(os.Getenv("DB_PORT"))
	//Postgres Repo
	db := database.NewDatabase(os.Getenv("DB_HOST"), port , os.Getenv("DB_DRIVER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	CaseRepository := repository.NewPostgresRepository(db)
	//Memory Repo
	//CaseRepository repository.CaseRepository = repository.NewMemoryRepository()
	//Other Params(setting up controller and service)
	caseService := service.NewCaseService(CaseRepository)
	caseController := controller.NewCaseCrontroller(caseService)
	r := gin.Default()

	r.GET("/cases", caseController.GetCases)
	r.GET("/case/:id", caseController.GetCase)
	r.POST("/cases", caseController.PostCase)
	r.PUT("/cases", caseController.PutCase)
	r.DELETE("/cases/:id", caseController.DeleteCase)

	return r
}
