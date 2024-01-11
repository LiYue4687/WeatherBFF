package service

import (
	"context"
	"fmt"
	"github.com/LiYue4687/WeatherBFF/grpcServer/repository"
	"github.com/LiYue4687/WeatherBFF/proto"
	"log"
)

type SearchService struct {
	*proto.UnimplementedCitySearchServiceServer
}

const PORT = "9001"

func (s *SearchService) Search(ctx context.Context, r *proto.CitySearchRequest) (*proto.CitySearchResponse, error) {
	items, err := repository.SearchCity(r)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("receive a request: ", r)
	return &proto.CitySearchResponse{Items: items}, nil
}
