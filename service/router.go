package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-spread-service/controller"
)

// Router ...
func Router(server *HTTPServer) *gin.Engine {
	version := "v0"
	eng := server.Engine

	//TODO
	//staticFS, err := fs.New()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//eng.NoRoute(func(ctx *gin.Context) {
	//	opened, err := staticFS.Open(ctx.Request.URL.Path)
	//	if ctx.Request.URL.Path == "/" || err != nil {
	//		opened, err = staticFS.Open("/index.html")
	//		if err != nil {
	//			ctx.AbortWithStatus(http.StatusNotFound)
	//			return
	//		}
	//	}
	//	ctx.Status(http.StatusOK)
	//	_, err = io.Copy(ctx.Writer, opened)
	//})

	spread := eng.Group("spread")

	spread.GET("/:sign/authorize", controller.AuthorizeNotify(version))

	////登录
	//g0.POST("login", LoginPOST(verV0))
	////组织注册
	//g0.POST("register", RegisterPOST(verV0))
	//
	//g0.POST("genesis", GenesisGet(verV0))
	return eng
}

func isInstalled() bool {
	return false
}

// AccessControlAllow ...
func AccessControlAllow(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("origin")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
		"Accept-Encoding, X-CSRF-Token, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
		return
	}
	ctx.Next()
}
