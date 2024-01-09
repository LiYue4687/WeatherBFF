package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weatherBFF/routes"
)

func main() {
	//创建一个默认的 Gin 引擎
	r := gin.Default()

	// 中间件：将 HTTP 重定向到 HTTPS
	r.Use(func(c *gin.Context) {
		if c.Request.TLS == nil {
			// 如果不是 HTTPS 请求，则重定向到 HTTPS
			url := "https://" + c.Request.Host + c.Request.RequestURI
			c.Redirect(http.StatusMovedPermanently, url)
			return
		}
		c.Next()
	})

	// 定义一个路由处理函数
	routes.GetWeatherForecast(r)
	routes.GetWeather24hour(r)

	// 启动 HTTPS 服务器
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", r)
	if err != nil {
		panic(err)
	}
}
