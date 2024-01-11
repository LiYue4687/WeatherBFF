package util

import (
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"math/rand"
	"strconv"
)

func SimulateWeather(reportTime, dayTemp, nightTemp int, weather string) []entity.Forecast24hourItem {
	hoursInDay := 24
	weatherData := make([]entity.Forecast24hourItem, hoursInDay)

	for hour := 0; hour < hoursInDay; hour++ {
		// 模拟温度波动，生成一个介于最低温和最高温之间的随机温度
		temperature := rand.Intn(dayTemp-nightTemp+1) + nightTemp
		weatherData[hour] =
			entity.Forecast24hourItem{
				Time:    strconv.Itoa((reportTime+hour)%24) + ":00",
				Temp:    strconv.Itoa(temperature),
				Weather: weather,
			}
	}

	return weatherData
}
