package initialize

import (
	"net/http"

	"WebWork/middleware"

	"WebWork/api"

	"WebWork/global"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.Use(middleware.Cors())
	Router.StaticFS("/file", http.Dir("./static/img"))
	global.G_LOG.Debug("use middleware cors")
	ApiGroup := Router.Group("")
	uploadRouter(ApiGroup)
	fiveRouter(ApiGroup)
	mitoRouter(ApiGroup)
	global.G_LOG.Debug("router register success")
	return Router
}

func uploadRouter(Router *gin.RouterGroup) {
	UploadRouter := Router.Group("upload")
	//.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UploadRouter.POST("/header", api.Upload)
		UploadRouter.POST("/img", api.Upload)
	}
}

func fiveRouter(Router *gin.RouterGroup) {
	Router.GET("/five/:key", api.GetFive)
}

func mitoRouter(Router *gin.RouterGroup) {
	Router.GET("/mito/:path/:index", api.SaveMito)
}
