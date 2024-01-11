package entity

type CitySearchResponse struct {
	Status string      `json:"status"`
	Cities []*CityItem `json:"cities"`
}

type CityItem struct {
	ProvinceName string `json:"provinceName"`
	CityName     string `json:"cityName"`
	CountyName   string `json:"countyName"`
	CityCode     string `json:"cityCode"`
	AdCode       string `json:"adCode"`
}

func (c CityItem) IsEqual(other CityItem) bool {
	return c.ProvinceName == other.ProvinceName &&
		c.CityName == other.CityName &&
		c.CountyName == other.CountyName &&
		c.CityCode == other.CityCode &&
		c.AdCode == other.AdCode
}
