package main

import (
	"deploy-tools/router"

	"deploy-tools/models"
	"deploy-tools/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitModels()
	r := gin.Default()
	httpport := models.GetConfigValue("httpport")
	router.StartRouter(r, utils.ParseInt(httpport))
}
