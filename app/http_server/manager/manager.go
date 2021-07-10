//+build manager

package manager

import (
	"github.com/gin-gonic/gin"
	"github/tqcenglish/word-english/web"
	"io/fs"
	"net/http"
)

func Enable(router *gin.Engine) {
	managerRoot, _ := fs.Sub(web.WWWStatic, "www")
	router.StaticFS("/www", http.FS(managerRoot))
}
