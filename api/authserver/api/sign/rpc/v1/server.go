package v1

import (
	"context"
	auth "github.com/zxm1124/component-base/pkg/auth/v1"
	"github.com/zxm1124/store-ming/api/authserver/api/sign/global"
	"github.com/zxm1124/store-ming/api/authserver/api/sign/rpc/v1/pb"
)

type SignServer struct{}

func (this *SignServer) SignToken(ctx context.Context, req *pb.SignReq) (*pb.SignResp, error) {
	instanceID := req.GetInstanceID()

	tokenString := auth.Sign(instanceID,
		global.AuthInfo.SignInfo.Audience,
		global.AuthInfo.SignInfo.Issuer,
		global.AuthInfo.SignInfo.Secret,
		global.AuthInfo.SignInfo.Timeout)

	return &pb.SignResp{
		TokenString: tokenString,
	}, nil
}
