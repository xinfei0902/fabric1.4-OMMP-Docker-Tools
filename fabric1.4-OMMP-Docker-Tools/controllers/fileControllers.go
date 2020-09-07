package controllers

import (
	"deploy-tools/models"

	"github.com/gin-gonic/gin"
	"github.com/monax-master/log"
)

//FileController 声明结构
type FileController struct {
}

// SetRouter 文件控制路由地址对应接口
func (con *FileController) SetRouter(e *gin.Engine) {
	g := e.Group("/command/file")
	g.POST("upload", con.UploadFile)
}

// UploadFile 上传文件
func (con *FileController) UploadFile(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			responseJSON(c, COMMAND_FUNCTION_UPLOAD)
		}
	}()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		code := CodeMessage(201, err)
		responseJSON(c, code)
		return
	}
	filename := header.Filename

	success, err := models.UploadExec(file, filename)
	if err != nil {
		code := CodeMessage(201, err)
		responseJSON(c, code)
		return
	}
	responseJSON(c, SUCCESSOK, "data", success)
	return
}
