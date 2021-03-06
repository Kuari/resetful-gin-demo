package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"resetful-gin-demo/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.Authorize())

	// 加载路由
	LoadUser(r)
	LoadTodo(r)
	LoadPing(r)

	return r
}
