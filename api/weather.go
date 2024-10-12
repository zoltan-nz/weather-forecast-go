package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// API Documentation: https://open-meteo.com/en/docs/geocoding-api
const OpenMeteoGeocodingApiUrl = "https://geocoding-api.open-meteo.com/v1/search"

type GeoResponse struct {
	Results []LatLong `json:"results"`
}

type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func FetchLatLong(city string) (float64, float64, error) {
	client := resty.New()
	queryParams := map[string]string{
		"name":     city,
		"count":    "1",
		"language": "en",
		"format":   "json",
	}
	response, err := client.R().SetQueryParams(queryParams).Get(OpenMeteoGeocodingApiUrl)

	if err != nil {
		return 0, 0, fmt.Errorf("error making request to OpenMeteoGeocodingApi: %w", err)
	}

	if response.StatusCode() != http.StatusOK {
		return 0, 0, fmt.Errorf("unexpected status code: %d and response: %s", response.StatusCode(), response.String())
	}

	var geoResponse GeoResponse
	if err := json.Unmarshal(response.Body(), &geoResponse); err != nil {
		return 0, 0, fmt.Errorf("error decoding response: %w", err)
	}

	if len(geoResponse.Results) == 0 {
		return 0, 0, fmt.Errorf("no results found for city: %s", city)
	}

	return geoResponse.Results[0].Latitude, geoResponse.Results[0].Longitude, nil
}
