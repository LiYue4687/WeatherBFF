package entity

// Weather24hourResponse response for 24hour forecast
type Weather24hourResponse struct {
	Status         string               `json:"status"`
	CityInfo       CityInfo             `json:"cityInfo"`
	Forecast24hour []Forecast24hourItem `json:"forecast24hour"`
}

type CityInfo struct {
	City       string `json:"city"`
	AdCode     string `json:"adCode"`
	Province   string `json:"province"`
	ReportTime string `json:"reportTime"`
}

type Forecast24hourItem struct {
	Time string `json:"time"`
	Temp string `json:"temp"`
}
