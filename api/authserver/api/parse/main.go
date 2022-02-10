package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	sViper "github.com/zxm1124/component-base/pkg/viper"
	controller "github.com/zxm1124/store-ming/api/authserver/api/parse/controller/v1"
	"github.com/zxm1124/store-ming/api/authserver/api/parse/global"
	"log"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET(global.AuthInfo.ParsePath, controller.OnAuth)

	err := router.Run(":" + strconv.Itoa(global.AuthInfo.ParseHttpPort))
	if err != nil {
		logrus.Panicf("Auth ParseToken Http Server listening failed, port at %d handlePath is %s",
			global.AuthInfo.ParseHttpPort,
			global.AuthInfo.ParsePath)
	}
}

// init函数
func init() {
	err := sViper.SetupSetting("auth", &global.AuthInfo)
	if err != nil {
		log.Panicf("init.setupSetting code: %v", err)
	}
}
