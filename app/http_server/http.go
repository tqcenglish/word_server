package http_server

import (
	"fmt"

	"github/tqcenglish/word-english/api"
	"github/tqcenglish/word-english/app/http_server/apidoc"
	"github/tqcenglish/word-english/app/http_server/manager"
	"github/tqcenglish/word-english/app/http_server/pprof"
	"github/tqcenglish/word-english/configs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateServer(port int) {
	// 创建路由
	router := gin.New()

	// 配置中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 设置静态目录
	router.Static("/word-static", fmt.Sprintf("%s/out", configs.BasicPath))

	//pprof
	pprof.Enable(router)

	// 设置 static
	manager.Enable(router)

	// apidoc
	apidoc.Enable(router)

	// API 路由
	api.AddRouter(router)

	addr := fmt.Sprintf("%s:%d", configs.ConfigGlobal.WebHost, configs.ConfigGlobal.WebPort)
	logrus.Infof("http addr: %s", addr)
	router.Run(addr)

}
