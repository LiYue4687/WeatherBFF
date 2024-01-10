package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"weatherBFF/bff/config"
	entity2 "weatherBFF/bff/entity"
	"weatherBFF/bff/util"
)

func GetWeatherForecast(city, extensions string) *entity2.WeatherForecastResponse {
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
	var res entity2.WeatherForecastResponse
	_ = json.Unmarshal(body, &res)
	return &res
}

func GetWeather24hour(city, extensions string) *entity2.Weather24hourResponse {
	res := GetWeatherForecast(city, extensions)
	var realRes entity2.Weather24hourResponse
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
