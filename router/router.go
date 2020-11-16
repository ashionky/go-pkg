package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-pkg/handler"
	"go-pkg/pkg/cfg"
	"go-pkg/pkg/middleware"
)

var config  = cfg.GetConfig()

func InitRouter() error {
	r := gin.Default()

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

	}

	pri := v1.Group("/pri", handler.Authorize)
	{
		// 需要登录才可访问的api
		pri.POST("/signout", handler.SignOut)
	}

	if err := r.Run(fmt.Sprintf(":%s", config.Server.Http_port)); err != nil {
		fmt.Print("运行失败：",err )
		return err
	}
	return nil
}
