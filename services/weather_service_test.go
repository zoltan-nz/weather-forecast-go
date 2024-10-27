package services

import (
	"fmt"
	"github.com/zoltan-nz/weather-forecast-go/models"
	"log"
	"os"
	"testing"
)

func TestNewWeatherService(t *testing.T) {
	logger := log.New(os.Stdout, "weather_service_test", log.LstdFlags)
	service := NewWeatherService(logger)

	if service == nil {
		t.Error("Expected weather service instance, got nil")
	}
}

func TestWeatherService_FetchWeather(t *testing.T) {
	logger := log.New(os.Stdout, "weather_service_test", log.LstdFlags)
	service := NewWeatherService(logger)

	tests := []struct {
		name    string
		latLong models.LatLong
		wantErr bool
	}{
		{
			name:    "Valid coordinates",
			latLong: models.LatLong{Lat: 43.70455, Long: -79.404625}, // Toronto
			wantErr: false,
		},
		{
			name:    "Invalid coordinates",
			latLong: models.LatLong{Lat: 91, Long: 181}, // Outside valid ranges
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weather, err := service.FetchWeather(tt.latLong)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && weather != nil {
				if len(weather.HourlyTemperatures) == 0 {
					t.Errorf("FetchWeather() hourly temperature is empty")
				}
			}

			if !tt.wantErr {
				validateWeatherData(t, weather)
			}
		})
	}
}

func validateWeatherData(t *testing.T, weather *models.WeatherData) {
	t.Helper()

	if weather == nil {
		t.Fatal("Expected weather data, got nil")
	}

	if len(weather.HourlyTemperatures) == 0 {
		t.Error("Expected non-empty hourly temperatures")
	}

	first := weather.HourlyTemperatures[0]
	if first.Time == "" {
		t.Error("Expected non-empty time string")
	}

	if first.Temperature < -100 || first.Temperature > 100 {
		t.Errorf("Expected temperature in range -100..100, got %f", first.Temperature)
	}
}

type mockWeatherService struct {
	shouldError bool
}

func (m *mockWeatherService) FetchWeather(latLong models.LatLong) (*models.WeatherData, error) {
	if m.shouldError {
		return nil, fmt.Errorf("error fetching weather data")
	}

	return &models.WeatherData{
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
	}, nil
}
