package models

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

const (
	COMMAND_UNPACK     = "unpack"     //解压命令
	COMMAND_START      = "start"      //脚本开始执行命令
	COMMAND_STATUS     = "status"     //执行获取节点状态脚本
	COMMAND_REMOVE     = "remove"     //移除
	COMMAND_CLEAR      = "clear"      //清除
	COMMAND_CUSTOMIZED = "customized" //自定义命令
)

//CommandExec 根据命令执行不同操作
func CommandExec(command string, args []string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			return
		}
	}()

	switch command {
	case COMMAND_UNPACK:
		//解压命令具体执行DO
		execCommand := args[0]
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", execCommand)
		} else {
			cmd = exec.Command("/bin/bash", "-c", execCommand)
		}
		stderr, _ := cmd.StderrPipe()
		if err := cmd.Start(); err != nil {
			log.Error("exec command", execCommand, "err info:", err)
			return "", err
		}

		bytes, err := ioutil.ReadAll(stderr)
		if err != nil {
			log.Error("ReadAll stdout fail:", err)
			return "", err
		}

		if err := cmd.Wait(); err != nil {
			log.Error("exec command wait:", err)
			return "", errors.New(string(bytes))
		}
		if !cmd.ProcessState.Success() {
			// 执行失败，返回错误信息
			return "", errors.New(string(bytes))
		}
		return "解压成功", nil
	case COMMAND_START:
		//执行命令
		execCommand := args[0]
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", execCommand)
		} else {
			cmd = exec.Command("/bin/bash", "-c", execCommand)
		}
		stderr, _ := cmd.StderrPipe()
		if err := cmd.Start(); err != nil {
			log.Error(" exec command", execCommand, "err info:", err)
			return "", err
		}
		var errBuf string
		bytes, err := ioutil.ReadAll(stderr)
		if err != nil {
			log.Error(" ReadAll stdout fail:", err)
			return "", err
		}
		errString := string(bytes)
		errInfoA := strings.Split(errString, "Error")
		if len(errInfoA) >= 2 {
			errBuf = errInfoA[1]
		} else {
			errBuf = errInfoA[0]
		}
		//cmd.Wait()
		if err := cmd.Wait(); err != nil {
			log.Error(" exec command wait:", err)
			return "", errors.New(errBuf)
		}
		if !cmd.ProcessState.Success() {
			// 执行失败，返回错误信息
			return "", errors.New(errBuf)
		}
		return "执行脚本成功", nil
	case COMMAND_STATUS:
		execCommand := args[0]
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", execCommand)
		} else {
			cmd = exec.Command("/bin/bash", "-c", execCommand)
		}
		err := cmd.Run()
		if err != nil {
			log.Error(" exec command", execCommand, "err info:", err)
			return "", err
		}
		return "执行脚本成功", nil
	case COMMAND_REMOVE:
		return "移除成功", nil
	case COMMAND_CLEAR:
		return "清理完毕", nil
	case COMMAND_CUSTOMIZED:
		return "", nil
	default:
		return "未知命令", nil
	}
}
