package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zoltan-nz/weather-forecast-go/config"
	"github.com/zoltan-nz/weather-forecast-go/models"
	"log"
	"net/http"
	"time"
)

type WeatherService interface {
	FetchWeather(latLong models.LatLong) (*models.WeatherData, error)
}

type weatherService struct {
	apiURL string
	client *resty.Client
	logger *log.Logger
}

func NewWeatherService(logger *log.Logger) WeatherService {
	return &weatherService{
		apiURL: config.OpenMeteoWeatherApiUrl,
		client: resty.New().SetTimeout(10 * time.Second),
		logger: logger,
	}
}

func (s *weatherService) FetchWeather(latLong models.LatLong) (*models.WeatherData, error) {
	s.logger.Printf("Fetching weather data for lat: %f, long: %f", latLong.Lat, latLong.Long)

	queryParams := map[string]string{
		"latitude":  fmt.Sprintf("%.6f", latLong.Lat),
		"longitude": fmt.Sprintf("%.6f", latLong.Long),
		"hourly":    "temperature_2m",
		"timezone":  "auto",
	}

	response, err := s.client.R().SetQueryParams(queryParams).Get(s.apiURL)

	if err != nil {
		return nil, fmt.Errorf("error making request to OpenMeteoWeatherApi: %w", err)
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d and response: %s", response.StatusCode(), response.String())
	}

	var weatherResponse models.WeatherResponse
	if err := json.Unmarshal(response.Body(), &weatherResponse); err != nil {
		return nil, fmt.Errorf("error parsing weather response: %w", err)
	}

	return deserializeWeatherData(&weatherResponse)
}

func deserializeWeatherData(response *models.WeatherResponse) (*models.WeatherData, error) {
	if response == nil {
		return nil, fmt.Errorf("response is nil")
	}

	timeLen := len(response.Hourly.Time)
	if timeLen != len(response.Hourly.Temperature) {
		return nil, fmt.Errorf("time and temperature slices have different lengths")
	}

	hourlyTemps := make([]models.HourlyTemperature, timeLen)

	for i := range response.Hourly.Time {
		hourlyTemps[i] = models.HourlyTemperature{
			Time:        response.Hourly.Time[i],
			Temperature: response.Hourly.Temperature[i],
		}
	}

	return &models.WeatherData{
		HourlyTemperatures: hourlyTemps,
	}, nil
}
