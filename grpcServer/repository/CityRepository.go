package repository

import (
	"database/sql"
	"log"
	proto "weatherBFF/grpcServer/service"

	_ "github.com/mattn/go-sqlite3"
)

func SearchCity(request *proto.CitySearchRequest) ([]*proto.CityItem, error) {

	// 打开database
	var db *sql.DB
	db, dbOpenErr := sql.Open("sqlite3",
		"/Users/lmc60019/GolandProjects/WeatherBFF/grpcServer/database/weather_database.db")
	if dbOpenErr != nil {
		log.Fatal(dbOpenErr)
	}
	// 执行SQLite查询
	rows, queryErr := db.Query(`
            SELECT city_info.province as provinceName, city_info.city as cityName, city_list.name as countyName, 
                   city_info.city_code as cityCode, city_list.ad_code as adCode
            FROM city_list
            INNER JOIN city_info ON city_list.city_code = city_info.city_code
            WHERE (city_list.name LIKE '%'||?||'%') OR (city_info.city LIKE '%'||?||'%');
            `, request.SearchValue, request.SearchValue)
	if queryErr != nil {
		log.Fatal(queryErr)
		return nil, queryErr
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("something unexpect happened while closing database")
		}
	}(db)

	// 处理查询结果并构建gRPC响应
	var items []*proto.CityItem
	for rows.Next() {
		var provinceName, cityName, countyName, cityCode, adCode string
		if err := rows.Scan(&provinceName, &cityName, &countyName, &cityCode, &adCode); err != nil {
			log.Fatal(err)
			return nil, err
		}
		// 根据需要处理结果，例如将其添加到响应中
		items = append(items, &proto.CityItem{ProvinceName: provinceName, CityName: cityName, CountyName: countyName,
			CityCode: cityCode, AdCode: adCode})
	}

	return items, nil
}
