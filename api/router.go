package api

import (
	"github/tqcenglish/word-english/configs"

	"github.com/gin-gonic/gin"
)

// AddRouter 添加路由
func AddRouter(router *gin.Engine) {
	router.Use(CORSMiddleware("*"))
	router.Use(JWTAuth(configs.ConfigGlobal.IdentityKey))

	wordEnglish := router.Group("/word-english")
	wordEnglish.GET("/word/:word", DetailWord)
	wordEnglish.GET("/word", GetAllWords)
	wordEnglish.POST("/play-mp3", PlayMP3)
}
