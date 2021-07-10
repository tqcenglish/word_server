//+build apidoc

package apidoc

import (
	"github.com/gin-gonic/gin"
	"github/tqcenglish/word-english/web"
	"io/fs"
	"net/http"
)

func Enable(router *gin.Engine) {
	ApiDocStatic, _ := fs.Sub(web.ApiDocStatic, "apidoc")
	router.StaticFS("/apidoc", http.FS(ApiDocStatic))
}
