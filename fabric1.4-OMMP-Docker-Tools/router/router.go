package router

import (
	"deploy-tools/controllers"
	"deploy-tools/utils"

	"github.com/gin-gonic/gin"
)

//StartRouter 配置路由
func StartRouter(e *gin.Engine, port int) {
	new(controllers.CommandController).SetRouter(e)
	new(controllers.FileController).SetRouter(e)
	e.Run(":" + utils.ParseString(port))
}
