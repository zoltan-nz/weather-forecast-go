package services

import (
	"fmt"
	"github.com/zoltan-nz/weather-forecast-go/models"
)

type MockWeatherService struct {
	ShouldError  bool
	CalledWith   models.LatLong
	CallCount    int
	MockResponse *models.WeatherData
}

func NewMockWeatherService() *MockWeatherService {
	return &MockWeatherService{
		MockResponse: &models.WeatherData{
			HourlyTemperatures: []models.HourlyTemperature{
				{
					Time:        "2021-06-01T00:00",
					Temperature: 20.5,
				},
				{
					Time:        "2021-06-01T01:00",
					Temperature: 19.8,
				},
			},
		},
	}
}

// FetchWeather is a mock implementation of the WeatherService interface
func (m *MockWeatherService) FetchWeather(latLong models.LatLong) (*models.WeatherData, error) {
	m.CalledWith = latLong
	m.CallCount++

	if m.ShouldError {
		return nil, fmt.Errorf("mock error")
	}

	return m.MockResponse, nil
}
