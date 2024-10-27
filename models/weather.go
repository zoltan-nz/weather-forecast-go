package models

type WeatherResponse struct {
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Hourly    HourlyResponse `json:"hourly"`
}

type HourlyResponse struct {
	Time        []string  `json:"time"`
	Temperature []float64 `json:"temperature_2m"`
}

type WeatherData struct {
	HourlyTemperatures []HourlyTemperature
}

type HourlyTemperature struct {
	Time        string
	Temperature float64
}
