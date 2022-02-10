package v1

import (
	"context"
	"github.com/sirupsen/logrus"
	auth "github.com/zxm1124/component-base/pkg/auth/v1"
	"github.com/zxm1124/store-ming/internal/authserver/meta"
	"github.com/zxm1124/store-ming/internal/authserver/rpc/v1/pb"
)

type SignServer struct{}

func (this *SignServer) SignToken(ctx context.Context, req *pb.SignReq) (*pb.SignResp, error) {
	instanceID := req.GetInstanceID()

	logrus.Infof("user: %s apply to sign a token", instanceID)

	tokenString := auth.Sign(instanceID,
		meta.AuthInfo.SignInfo.Audience,
		meta.AuthInfo.SignInfo.Issuer,
		meta.AuthInfo.SignInfo.Secret,
		meta.AuthInfo.SignInfo.Timeout)

	return &pb.SignResp{
		TokenString: tokenString,
	}, nil
}
