package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	myLog "github.com/zxm1124/component-base/log"
	sViper "github.com/zxm1124/component-base/pkg/viper"
	controller "github.com/zxm1124/store-ming/internal/authserver/controller/v1"
	"github.com/zxm1124/store-ming/internal/authserver/meta"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET(meta.AuthInfo.ParsePath, controller.OnAuth)

	err := router.Run(":" + strconv.Itoa(meta.AuthInfo.ParseHttpPort))
	if err != nil {
		logrus.Panicf("Auth ParseToken Http Server listening failed, port at %d handlePath is %s",
			meta.AuthInfo.ParseHttpPort,
			meta.AuthInfo.ParsePath)
	}
}

// init函数
func init() {

	err := sViper.SetupSetting("auth", &meta.AuthInfo)
	if err != nil {
		logrus.Panicf("init.setupSetting code: %v", err)
	}

	// 设置日志
	err = sViper.SetupSetting("ParseLog", &meta.LogInfo)
	if err != nil {
		logrus.Panicf("init.setupSetting code: %v", err)
	}

	myLog.SetLog(
		meta.LogInfo.Path,
		meta.LogInfo.Level,
		meta.LogInfo.Formatter,
		meta.LogInfo.Output)

}
