package v1

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	auth "github.com/zxm1124/component-base/pkg/auth/v1"
	"github.com/zxm1124/component-base/pkg/code"
	"github.com/zxm1124/store-ming/api/authserver/api/parse/global"
)

// OnAuth 校验用户token
func OnAuth(c *gin.Context) {

	// 用户没有token
	tokenString := c.GetHeader(("X-Authentication"))
	if tokenString == "" {
		err := code.ErrTokenInvalid

		resp := gin.H{
			"code": err.Code(),
			"msg":  err.Msg(),
		}

		c.JSON(err.HttpStatus(), resp)

		log.Info("token is nil")

		return
	}

	secret := global.AuthInfo.SignInfo.Secret

	if ok, _ := auth.ParseToken(tokenString, secret); !ok {
		err := code.ErrTokenInvalid

		resp := gin.H{
			"code": err.Code(),
			"msg":  err.Msg(),
		}

		c.JSON(err.HttpStatus(), resp)

		log.Info("parse token failed, token is invalid")

		return
	}

	// 用户有token
	if ok, claims := auth.ParseToken(tokenString, secret); ok {
		err := code.ErrSuccess

		resp := gin.H{
			"code": err.Code(),
			"msg":  err.Msg(),
		}

		c.JSON(err.HttpStatus(), resp)

		log.WithFields(log.Fields{
			"instanceID": claims.InstanceID,
			"expired":    claims.ExpiresAt,
		}).Info("parse token succeeded")
	}
}
