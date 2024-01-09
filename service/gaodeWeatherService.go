package service

import (
	"weatherBFF/entity"
	"weatherBFF/repository"
)

func GetWeatherForecast(city, extensions string) *entity.WeatherForecastResponse {
	return repository.GetWeatherForecast(city, extensions)
}

func GetWeather24hour(city, extensions string) interface{} {
	return repository.GetWeather24hour(city, extensions)
}
