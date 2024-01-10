package repository

import (
	"fmt"
	"testing"
)

// 白盒测试

func TestGetWeatherForecast(t *testing.T) {
	testRes := GetWeatherForecast("110101", "all")
	fmt.Println(testRes)
	if testRes.Status != "1" {
		t.Errorf("运行出错，status != 1")
	}
	if testRes.Forecasts[0].AdCode != "110101" {
		t.Errorf("获取结果adcode与预期不符")
	}
	if len(testRes.Forecasts[0].Casts) != 4 {
		t.Errorf("获取预测数量与预期不符")
	}

}

func TestGetWeather24hour(t *testing.T) {
	testRes := GetWeather24hour("110101", "all")
	fmt.Println(testRes)
	if testRes.Status != "1" {
		t.Errorf("运行出错，status != 1")
	}
	if testRes.CityInfo.AdCode != "110101" {
		t.Errorf("获取结果adcode与预期不符")
	}
	if len(testRes.Forecast24hour) != 24 {
		t.Errorf("获取预测数量与预期不符")
	}

}
