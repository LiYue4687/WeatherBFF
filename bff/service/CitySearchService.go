package service

import (
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"github.com/LiYue4687/WeatherBFF/bff/repository"
)

func SearchCity(searchValue string) []*entity.CityItem {
	return repository.SearchCity(searchValue)
}
