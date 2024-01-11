package repository

import (
	"github.com/LiYue4687/WeatherBFF/bff/entity"
	"github.com/LiYue4687/WeatherBFF/bff/repository"
	"testing"
)

// 黑盒测试
func TestCitySearch(t *testing.T) {
	cases := []struct {
		Name     string
		Expected entity.CitySearchResponse
	}{
		{
			Name: "金州",
			Expected: entity.CitySearchResponse{
				Status: "1", Cities: []*entity.CityItem{
					{"辽宁省", "大连市", "金州区", "0411", "210213"},
				},
			},
		},
	}

	for _, c := range cases {
		testRes := repository.SearchCity(c.Name)
		if !isEqual(testRes, c.Expected.Cities) {
			t.Fatalf("TestCitySearch get wrong data")
		}
	}
}

func isEqual(l1, l2 []*entity.CityItem) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, c := range l1 {
		if !c.IsEqual(*l2[i]) {
			return false
		}
	}
	return true
}
