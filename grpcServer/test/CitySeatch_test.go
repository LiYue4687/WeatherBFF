package test

import (
	"fmt"
	"github.com/LiYue4687/WeatherBFF/grpcServer/repository"
	"github.com/LiYue4687/WeatherBFF/proto"
	"testing"
)

func TestSearchCity(t *testing.T) {
	city, err := repository.SearchCity(&proto.CitySearchRequest{
		SearchValue: "大连",
	})
	if err != nil {
		return
	}
	fmt.Println(city)
}
