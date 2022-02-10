package main

import (
	"context"
	"fmt"
	"github.com/zxm1124/store-ming/api/authserver/api/sign/rpc/v1/pb"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	//客户端连接服务器
	conn, err := grpc.Dial(":8181", grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络异常", err)
	}
	defer conn.Close()

	//获取grpc句柄
	c := pb.NewSignServerClient(conn)

	resp, err := c.SignToken(context.TODO(), &pb.SignReq{InstanceID: "user-nkjj12"})
	if err != nil {
		fmt.Println("code:", err)
	}
	fmt.Println("resp.tokenString:", resp.TokenString)
}
