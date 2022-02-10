package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	controller "store-ming/api/authserver/api/parse/controller/v1"
	"store-ming/api/authserver/api/parse/global"
	sViper "store-ming/component-base/pkg/viper"
	"strconv"
)

func main() {
	http.HandleFunc(global.AuthInfo.ParsePath, controller.OnAuth)

	log.Infof("Auth ParseToken Http Server listening port at %d, the handle api is '%s'",
		global.AuthInfo.ParseHttpPort,
		global.AuthInfo.ParsePath)

	_ = http.ListenAndServe(":"+strconv.Itoa(global.AuthInfo.ParseHttpPort), nil)
}

// init函数
func init() {
	err := sViper.SetupSetting("auth", &global.AuthInfo)
	if err != nil {
		log.Panicf("init.setupSetting code: %v", err)
	}
}
