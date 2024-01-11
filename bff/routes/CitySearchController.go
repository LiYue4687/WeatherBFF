package routes

import (
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"github.com/LiYue4687/WeatherBFF/bff/service"
	"github.com/gin-gonic/gin"
)

func GetSearchCity(r *gin.Engine) {
	r.GET("/GetSearchCity", func(c *gin.Context) {
		// 使用 c.Param 获取 URL 中的参数
		searchValue := c.Query("searchValue")

		// 调用 service.GetWeather 函数并返回JSON响应
		result := service.SearchCity(searchValue)
		// 判断 result 是否为 nil
		if result == nil {
			c.JSON(500, gin.H{"error": "unexpected error"})
			return
		}

		c.JSON(200, entity.CitySearchResponse{Status: "1", Cities: result})
	})
}
