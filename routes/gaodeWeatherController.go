package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"weatherBFF/service"
)

func GetWeatherForecast(r *gin.Engine) {
	r.GET("/GetWeatherForecast", func(c *gin.Context) {
		// 使用 c.Param 获取 URL 中的参数
		city := c.Query("city")
		extensions := c.Query("extensions")
		fmt.Println(city, extensions)

		// 调用 service.GetWeather 函数并返回JSON响应
		result := service.GetWeatherForecast(city, extensions)
		// 判断 result 是否为 nil
		if result == nil {
			c.JSON(500, gin.H{"error": "unexpected error"})
			return
		}

		// 根据 result 的内部变量确定返回值
		switch result.Status {
		case "1":
			c.JSON(200, result)
		case "0":
			c.JSON(404, gin.H{"error": "weather data not found"})
		}
	})
}

func GetWeather24hour(r *gin.Engine) {
	r.POST("/GetWeather24hour", func(c *gin.Context) {
		city := c.PostForm("city")
		extensions := c.PostForm("extensions")

		result := service.GetWeather24hour(city, extensions)
		// 判断 result 是否为 nil
		if result == nil {
			c.JSON(500, gin.H{"error": "unexpected error"})
			return
		}
	})
}
