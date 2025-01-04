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
	Router.StaticFS("/img", http.Dir("./static/img"))
	Router.StaticFS("/file", http.Dir("./static"))
	Router.Static("/css", "./css")
	Router.LoadHTMLGlob("./html/*")
	global.G_LOG.Debug("use middleware cors")
	ApiGroup := Router.Group("")
	uploadRouter(ApiGroup)
	fiveRouter(ApiGroup)
	getFiles(ApiGroup)
	mitoRouter(ApiGroup)
	htmlRouter(ApiGroup)
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
func htmlRouter(Router *gin.RouterGroup) {
	htmlRouter := Router.Group("html")
	{
		htmlRouter.GET("/upload", func(c *gin.Context) {
			c.HTML(200, "upload.html", nil)
		})
		htmlRouter.GET("/imgs", func(c *gin.Context) {
			c.HTML(200, "imgList.html", nil)
		})
	}

}

func getFiles(Router *gin.RouterGroup) {
	Router.GET("/fileImgs", api.AllImg)
}
