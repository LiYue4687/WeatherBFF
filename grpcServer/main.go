package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"weatherBFF/grpcServer/repository"
	proto "weatherBFF/grpcServer/service"

	_ "github.com/mattn/go-sqlite3"
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
