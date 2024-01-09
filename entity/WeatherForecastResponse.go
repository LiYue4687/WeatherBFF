package entity

// WeatherForecastResponse response for 3day forecast
type WeatherForecastResponse struct {
	Status    string     `json:"status"`
	Forecasts []Forecast `json:"forecasts"`
}

type Forecast struct {
	City       string         `json:"city"`
	AdCode     string         `json:"adCode"`
	Province   string         `json:"province"`
	ReportTime string         `json:"reportTime"`
	Casts      []ForecastItem `json:"casts"`
}

type ForecastItem struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	DayWeather   string `json:"dayWeather"`
	NightWeather string `json:"nightWeather"`
	DayTemp      string `json:"dayTemp"`
	NightTemp    string `json:"nightTemp"`
	DayWind      string `json:"dayWind"`
	NightWind    string `json:"nightWind"`
	DayPower     string `json:"dayPower"`
	NightPower   string `json:"nightPower"`
}
