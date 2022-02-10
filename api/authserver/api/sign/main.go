package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	global "store-ming/api/authserver/api/sign/global"
	signRpcV1 "store-ming/api/authserver/api/sign/rpc/v1"
	"store-ming/api/authserver/api/sign/rpc/v1/pb"
	sViper "github.com/zxm1124/component-base/pkg/viper"
	"strconv"
)

func main() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(global.AuthInfo.SignRpcPort))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Panicf("net.Listen failed: %v", err)
	}

	// 创建grpc服务
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterSignServerServer(s, &signRpcV1.SignServer{})

	//等待网络连接
	logrus.Infof("Auth SignToken Rpc Server listening port at %d",
		global.AuthInfo.SignRpcPort)

	err = s.Serve(listener)
	if err != nil {
		logrus.Panicf("grpc.Server failed: %v", err)
	}
}

// init函数
func init() {
	err := sViper.SetupSetting("auth", &global.AuthInfo)
	if err != nil {
		log.Panicf("init.setupSetting code: %v", err)
	}
}
