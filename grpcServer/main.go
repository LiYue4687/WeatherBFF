package main

import (
	"github.com/LiYue4687/WeatherBFF/grpcServer/service"
	"github.com/LiYue4687/WeatherBFF/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterCitySearchServiceServer(server, &service.SearchService{})

	lis, listenErr := net.Listen("tcp", ":"+service.PORT)
	if listenErr != nil {
		log.Fatalf("net.Listen err: %v", listenErr)
	}

	serverStartErr := server.Serve(lis)
	if serverStartErr != nil {
		return
	}
}
