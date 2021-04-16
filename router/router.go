package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-pkg/handler"
	"go-pkg/pkg/middleware"
	"go-pkg/pkg/websocket"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//websocket链接请求
	r.GET("/ws", websocket.WS)

	//静态文件
	r.Static("/static", "static")

	// swagger文档
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 跨域
	r.Use(middleware.Cors())

	v1 := r.Group("/v1")
	v1.Use(middleware.LogRequestMiddleware)
	v1.Use(middleware.LogResponseMiddleware)
	pub := v1.Group("/pub")
	{
		// 不需要登录即可访问的api
		pub.POST("/signin", handler.SignIn)
		//文件上传
		pub.POST("/uploadfile", handler.UploadFile)

	}

	pri := v1.Group("/pri", handler.Authorize)
	{
		// 需要登录才可访问的api
		pri.POST("/signout", handler.SignOut)
	}

	return r
}
