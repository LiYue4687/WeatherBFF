package repository_test

import (
	"fmt"
	"testing"
	"weatherBFF/repository"
)

// 黑盒测试
func TestGetWeatherForecast(t *testing.T) {
	testRes := repository.GetWeatherForecast("110101", "all")
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

// 白盒测试
func TestGetWeather24hour(t *testing.T) {
	testRes := repository.GetWeather24hour("110101", "all")
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

// 基线测试，测试运行效率
func BenchmarkWeather(b *testing.B) {
	testRes1 := repository.GetWeatherForecast("110101", "all")
	fmt.Println(testRes1)
	testRes2 := repository.GetWeather24hour("110101", "all")
	fmt.Println(testRes2)
}
