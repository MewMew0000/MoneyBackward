package setup

import (

	"github.com/gin-gonic/gin"

)

func InitRouter () *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("",func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	return router
}
