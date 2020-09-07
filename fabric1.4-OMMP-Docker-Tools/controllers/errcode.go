package controllers

import "fmt"

//Code 组装Code
func Code(code int, message string) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["message"] = message
	return m
}

var (
	//ERRAPINOTDONE 接口未完成
	ERRAPINOTDONE = Code(999, "接口未完成")
	//SUCCESSOK 正常
	SUCCESSOK               = Code(200, "OK")
	ERROR                   = Code(201, "错误")
	COMMAND_FUNCTION_ERROR  = Code(210, "因不可抗拒因素command执行失败")
	PARAM_COMMAND_EMPTY     = Code(211, "传进参数Command为空")
	PARAMMARSHALERROR       = Code(212, "传进参数解析错误")
	COMMAND_EXEC_COURSE     = Code(213, "脚本执行过程出错")
	COMMAND_FUNCTION_UPLOAD = Code(230, "因不可抗拒因素Upload执行失败")

	COMMAND_UPLOAD_FILE        = Code(250, "http 请求文件失败")
	COMMAND_UPLOAD_EXEC_COURSE = Code(251, "文件上传过程出错")
	PARAMERROR                 = Code(300, "传递参数错误")
)

func CodeMessage(code int, message error) map[string]interface{} {
	m := make(map[string]interface{})
	m["code"] = code
	m["message"] = fmt.Sprintf("exec fail:%s", message)
	return m
}
