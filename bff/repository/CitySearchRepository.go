package repository

import (
	"context"
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"github.com/LiYue4687/WeatherBFF/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func SearchCity(searchValue string) []*entity.CityItem {
	// 连接grpc服务器
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error happpened while closing connection: %v", err)
		}
	}(conn)

	// 初始化Greeter服务客户端
	c := proto.NewCitySearchServiceClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	r, err := c.Search(ctx, &proto.CitySearchRequest{SearchValue: searchValue})
	if err != nil {
		log.Fatalf("search fail: %v", err)
	}

	Cities := make([]*entity.CityItem, len(r.Items))
	for i, cityItem := range r.Items {
		Cities[i] = &entity.CityItem{
			ProvinceName: cityItem.ProvinceName,
			CityName:     cityItem.CityName,
			CountyName:   cityItem.CountyName,
			CityCode:     cityItem.CityCode,
			AdCode:       cityItem.AdCode,
		}
	}

	return Cities
}
