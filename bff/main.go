package main

import (
	"github.com/LiYue4687/WeatherBFF/bff/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的 Gin 引擎
	r := gin.Default()

	// 定义一个路由处理函数
	routes.GetWeatherForecast(r)
	routes.GetWeather24hour(r)
	routes.GetSearchCity(r)

	// 启动 HTTPS 服务器
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", r)
	if err != nil {
		panic(err)
	}
}
