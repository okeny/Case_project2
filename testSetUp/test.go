package testsetup

import "github.com/gin-gonic/gin"

func SetupTesting() *gin.Engine {

	r := gin.Default()
	gin.SetMode(gin.TestMode)

	return r
}
