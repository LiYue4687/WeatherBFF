package main

import (
	"github.com/gin-gonic/gin"
	"weatherBFF/routes"
)

func main() {
	//创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义一个路由处理函数
	routes.GetWeatherForecast(r)
	routes.GetWeather24hour(r)

	// 运行服务器，默认监听在 :8080
	err := r.Run()
	if err != nil {
		return
	}
}
