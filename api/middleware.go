package api

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var ignoreAuthPathMap = map[string]bool{
	"/ginapi/tts/xf/createTTS": true,
}

//JWTAuth 认证
func JWTAuth(apiKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		logrus.Tracef("token: %s", auth)

		if _, exist := ignoreAuthPathMap[ctx.Request.URL.Path]; exist {
			ctx.Next()
			return
		}

		if apiKey == "" {
			ctx.Next()
			return
		}

		if len(auth) < 7 {
			goto failure
		}

		{
			// Bearer xxx
			tokenString := strings.TrimSpace(auth[7:])

			// 添加 token 认证
			token, _ := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(apiKey), nil
			})

			if _, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
				ctx.Next()
				return
			}
		}

	failure:
		logrus.Warnf("jwt auth failure")
		ctx.AbortWithStatus(403)
	}
}

//CORSMiddleware cors
func CORSMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Request.Header.Del("Origin")
		c.Next()
	}
}
