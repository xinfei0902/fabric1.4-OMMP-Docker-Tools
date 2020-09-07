package models

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
// COMMAND_UNPACK     = "unpack"     //解压命令
// COMMAND_START      = "start"      //脚本开始执行命令
// COMMAND_REMOVE     = "remove"     //移除
// COMMAND_CLEAR      = "clear"      //清除
// COMMAND_CUSTOMIZED = "customized" //自定义命令
)

//CommandExec 根据命令执行不同操作
func UploadExec(file multipart.File, fileName string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			return
		}
	}()
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("workpath", workPath)
	fileSavePath := filepath.Join(workPath, fileName)
	_, err = os.Stat(fileSavePath)
	if err == nil {
		var cmd *exec.Cmd
		execCommand := fmt.Sprintf("rm -rf %s ", fileSavePath)
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", execCommand)
		} else {
			cmd = exec.Command("/bin/bash", "-c", execCommand)
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Error(" exec Stdout", execCommand, "err info:", err)
			return "", err
		}
		if err := cmd.Start(); err != nil {
			log.Error(" exec command", execCommand, "err info:", err)
			return "", err
		}
		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			log.Error(" ReadAll stdout fail:", err)
			return "", err
		}
		//fmt.Println("check stout info", string(bytes))
		if err := cmd.Wait(); err != nil {
			log.Error(" exec command wait:", err)
			return "", err
		}
		if !cmd.ProcessState.Success() {
			// 执行失败，返回错误信息
			return "", errors.New(string(bytes))
		}
	}
	out, err := os.Create(fileSavePath)
	if err != nil {
		log.Error("create save file", fileName, "err info:", err)
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Error("copy file info", fileName, "err info:", err)
		return "", err
	}
	return "文件成功上传", nil
}
