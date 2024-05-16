package gin

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/images")

	return r
}
