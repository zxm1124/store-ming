package main

import (
	"github.com/sirupsen/logrus"
	sViper "github.com/zxm1124/component-base/pkg/viper"
	signRpcV1 "github.com/zxm1124/store-ming/internal/authserver/api/sign/rpc/v1"
	"github.com/zxm1124/store-ming/internal/authserver/api/sign/rpc/v1/pb"
	"github.com/zxm1124/store-ming/internal/authserver/meta"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func main() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(meta.AuthInfo.SignRpcPort))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Panicf("net.Listen failed: %v", err)
	}

	// 创建grpc服务
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterSignServerServer(s, &signRpcV1.SignServer{})

	//等待网络连接
	logrus.Infof("Auth SignToken Rpc Server listening port at %d",
		meta.AuthInfo.SignRpcPort)

	err = s.Serve(listener)
	if err != nil {
		logrus.Panicf("grpc.Server failed: %v", err)
	}
}

// init函数
func init() {
	err := sViper.SetupSetting("auth", &meta.AuthInfo)
	if err != nil {
		log.Panicf("init.setupSetting code: %v", err)
	}
}
