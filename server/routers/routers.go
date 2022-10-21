package routers

import (
	"dcadmin/api"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	//跨域访问
	mwCORS := cors.New(cors.Config{
		// 允许跨域请求网站
		AllowOrigins: []string{"*"},
		// 允许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		// 允许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		// 显示请求的表头
		ExposeHeaders: []string{"Content-Type"},
		// 凭证共享， 确定共享
		AllowCredentials: true,
		// 允许跨域的原点网站
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})

	e.Use(mwCORS)

	e.GET("/users", api.ListUser)
	e.POST("/add", api.AddUser)
	e.GET("/verify", api.VerifyUser)
	e.Run()
}
