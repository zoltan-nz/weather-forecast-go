package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// OpenMeteoGeocodingApiUrl
// API Documentation: https://open-meteo.com/en/docs/geocoding-api
const OpenMeteoGeocodingApiUrl = "https://geocoding-api.open-meteo.com/v1/search"

type GeoResponse struct {
	Results []LatLongResponse `json:"results"`
}

type LatLongResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LatLong struct {
	Lat  float64
	Long float64
}

func FetchLatLong(city string) (LatLong, error) {
	client := resty.New()
	queryParams := map[string]string{
		"name":     city,
		"count":    "1",
		"language": "en",
		"format":   "json",
	}
	response, err := client.R().SetQueryParams(queryParams).Get(OpenMeteoGeocodingApiUrl)

	if err != nil {
		return LatLong{}, fmt.Errorf("error making request to OpenMeteoGeocodingApi: %w", err)
	}

	if response.StatusCode() != http.StatusOK {
		return LatLong{}, fmt.Errorf("unexpected status code: %d and response: %s", response.StatusCode(), response.String())
	}

	var geoResponse GeoResponse
	if err := json.Unmarshal(response.Body(), &geoResponse); err != nil {
		return LatLong{}, fmt.Errorf("error decoding response: %w", err)
	}

	if len(geoResponse.Results) == 0 {
		return LatLong{}, fmt.Errorf("no results found for city: %s", city)
	}

	return LatLong{geoResponse.Results[0].Latitude, geoResponse.Results[0].Longitude}, nil
}
