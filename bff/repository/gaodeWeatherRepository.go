package repository

import (
	"encoding/json"
	"fmt"
	"github.com/LiYue4687/WeatherBFF/bff/config"
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"github.com/LiYue4687/WeatherBFF/bff/util"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetWeatherForecast(city, extensions string) *entity.WeatherForecastResponse {
	params := url.Values{}
	Url, err := url.Parse(config.BaseUrl + "weatherInfo")
	if err != nil {
		return nil
	}
	params.Set("key", config.Key)
	params.Set("city", city)
	params.Set("extensions", extensions)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error when closing response body")
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	var res entity.WeatherForecastResponse
	_ = json.Unmarshal(body, &res)
	return &res
}

func GetWeather24hour(city, extensions string) *entity.Weather24hourResponse {
	res := GetWeatherForecast(city, extensions)
	var realRes entity.Weather24hourResponse
	realRes.Status = res.Status
	realRes.CityInfo.City = res.Forecasts[0].City
	realRes.CityInfo.Province = res.Forecasts[0].Province
	realRes.CityInfo.AdCode = res.Forecasts[0].AdCode
	realRes.CityInfo.ReportTime = res.Forecasts[0].ReportTime

	parsedTime, _ := time.Parse("2006-01-02 15:04:05", res.Forecasts[0].ReportTime)
	fmt.Println(parsedTime, parsedTime.Hour())
	dayTemp, _ := strconv.Atoi(res.Forecasts[0].Casts[0].DayTemp)
	nightTemp, _ := strconv.Atoi(res.Forecasts[0].Casts[0].NightTemp)
	realRes.Forecast24hour = util.SimulateWeather(parsedTime.Hour(), dayTemp, nightTemp)

	return &realRes
}
