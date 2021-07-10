//+build pprof

package pprof

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Enable(router *gin.Engine) {
	pprof.Register(router)
}
