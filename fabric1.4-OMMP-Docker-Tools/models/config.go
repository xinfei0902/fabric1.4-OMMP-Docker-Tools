package models

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	log    = logrus.New()
	config *goconfig.ConfigFile
)

//启动服务初始化
func InitModels() {
	loadConfig() //初始化位置文件
	fmt.Println("init config done!")
	log.AddHook(newLfsHook())
	fmt.Println("init log engine done!")
	gin.SetMode(GetConfigValue("mode"))
	fmt.Printf("set system mode to %s.", GetConfigValue("mode"))
}

//初始换加载配置文件路径
func loadConfig() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	fmt.Println(appConfigPath)
	config, err = goconfig.LoadConfigFile(appConfigPath)
}

//根据key值获取字段
func GetConfigValue(key string) string {
	runmode, err := config.GetValue(goconfig.DEFAULT_SECTION, "runmode")
	if err != nil {
		runmode = goconfig.DEFAULT_SECTION
	}
	value, _ := config.GetValue(runmode, key)
	if err != nil {
		return ""
	}
	return value
}

func newLfsHook() logrus.Hook {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logName := filepath.Join(workPath, "logs", GetConfigValue("appname")+".log")
	writer, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间,这里设置为一月分割一次
		rotatelogs.WithRotationTime(30*24*time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(30*24),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	return lfsHook
}
