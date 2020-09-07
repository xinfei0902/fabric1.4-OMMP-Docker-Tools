package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ParseRequestBody 解析gin的POST Body
func parseRequestBody(c *gin.Context) (request map[string]interface{}, err error) {
	var data []byte

	data, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("request:%s", data)
	err = json.Unmarshal(data, &request)
	if err != nil {
		return nil, err
	}
	return request, err
}

func responseJSON(c *gin.Context, code map[string]interface{}, args ...interface{}) {
	body := make(map[string]interface{})
	body["code"] = code
	for i := 0; i < len(args); i += 2 {

		switch args[i].(type) {
		case string:
			body[args[i].(string)] = args[i+1]
			break
		}
	}
	c.JSON(http.StatusOK, body)
}
