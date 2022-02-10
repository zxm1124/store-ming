package v1

import (
	log "github.com/sirupsen/logrus"
	auth "github.com/zxm1124/component-base/pkg/auth/v1"
	"net/http"
	"store-ming/api/authserver/api/parse/global"
)

// OnAuth 校验用户token
func OnAuth(w http.ResponseWriter, req *http.Request) {

	// 用户没有token
	tokenString := req.Header.Get("X-Authentication")
	if tokenString == "" {
		// 需要进行登录跳转 返回302
		_, _ = w.Write([]byte("The user has not logged in yet, and is jumping to the login page"))
		w.WriteHeader(http.StatusFound)

		log.Info("token is nil")

		return
	}

	secret := global.AuthInfo.SignInfo.Secret

	if ok, _ := auth.ParseToken(tokenString, secret); !ok {
		// 需要进行登录跳转 返回302
		_, _ = w.Write([]byte("The user has not logged in yet, and is jumping to the login page"))
		w.WriteHeader(http.StatusFound)

		log.Info("parse token failed, token is invalid")

		return
	}

	// 用户有token
	if ok, claims := auth.ParseToken(tokenString, secret); ok {
		_, _ = w.Write([]byte("authenticated success"))
		// 校验成功 返回200
		w.WriteHeader(http.StatusOK)

		log.WithFields(log.Fields{
			"instanceID": claims.InstanceID,
			"expired":    claims.ExpiresAt,
		}).Info("parse token succeeded")
	}
}
