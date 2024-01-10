package service

import (
	entity2 "weatherBFF/bff/entity"
	"weatherBFF/bff/repository"
)

func GetWeatherForecast(city, extensions string) *entity2.WeatherForecastResponse {
	return repository.GetWeatherForecast(city, extensions)
}

func GetWeather24hour(city, extensions string) *entity2.Weather24hourResponse {
	return repository.GetWeather24hour(city, extensions)
}
