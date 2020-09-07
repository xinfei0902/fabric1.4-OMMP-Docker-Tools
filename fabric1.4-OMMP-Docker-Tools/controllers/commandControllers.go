package controllers

import (
	"deploy-tools/models"
	"deploy-tools/utils"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/monax-master/log"
)

type CommandController struct {
}

func (con *CommandController) SetRouter(e *gin.Engine) {
	g := e.Group("/command")
	g.POST("exec", con.CommandExec)
}

//执行命令
func (con *CommandController) CommandExec(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			responseJSON(c, COMMAND_FUNCTION_ERROR)
		}
	}()

	request, err := parseRequestBody(c)
	if err != nil {
		code := CodeMessage(201, err)
		responseJSON(c, code)
		return
	}
	var command = utils.ParseString(request["command"])
	if len(command) == 0 {
		code := CodeMessage(201, err)
		responseJSON(c, code)
		return
	}
	var args []string
	argsL, err := json.Marshal(request["args"])
	if err != nil {
		log.Error("models CommandExec function incoming Param Marshal error")
		code := CodeMessage(201, err)
		responseJSON(c, code)
		return
	}
	json.Unmarshal(argsL, &args)
	// fmt.Println("command function command:", command)
	// fmt.Println("command function args:", args[0])
	// fmt.Println("command function args:", args[1])

	success, err := models.CommandExec(command, args)
	if err != nil {
		code := CodeMessage(201, err)
		responseJSON(c, code)
		return
	}
	responseJSON(c, SUCCESSOK, "data", success)
	return
}
