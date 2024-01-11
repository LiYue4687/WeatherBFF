package main

import (
	"context"
	"github.com/LiYue4687/WeatherBFF/grpcServer/repository"
	"github.com/LiYue4687/WeatherBFF/proto"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchService struct {
	*proto.UnimplementedCitySearchServiceServer
}

func (s *SearchService) Search(ctx context.Context, r *proto.CitySearchRequest) (*proto.CitySearchResponse, error) {
	items, err := repository.SearchCity(r)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &proto.CitySearchResponse{Items: items}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	proto.RegisterCitySearchServiceServer(server, &SearchService{})

	lis, listenErr := net.Listen("tcp", ":"+PORT)
	if listenErr != nil {
		log.Fatalf("net.Listen err: %v", listenErr)
	}

	serverStartErr := server.Serve(lis)
	if serverStartErr != nil {
		return
	}
}
