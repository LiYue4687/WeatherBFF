package service

import (
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"github.com/LiYue4687/WeatherBFF/bff/repository"
)

func GetWeatherForecast(city, extensions string) *entity.WeatherForecastResponse {
	return repository.GetWeatherForecast(city, extensions)
}

func GetWeather24hour(city, extensions string) *entity.Weather24hourResponse {
	return repository.GetWeather24hour(city, extensions)
}
