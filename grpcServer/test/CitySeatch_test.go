package test

import (
	"fmt"
	"testing"
	"weatherBFF/grpcServer/repository"
	"weatherBFF/proto"
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
