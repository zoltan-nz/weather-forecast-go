package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zoltan-nz/weather-forecast-go/config"
	"github.com/zoltan-nz/weather-forecast-go/models"
	"net/http"
)

func FetchLatLong(city string) (models.LatLong, error) {
	client := resty.New()
	queryParams := map[string]string{
		"name":  city,
		"count": "1",
	}
	response, err := client.R().SetQueryParams(queryParams).Get(config.OpenMeteoGeocodingApiUrl)

	if err != nil {
		return models.LatLong{}, fmt.Errorf("error making request to OpenMeteoGeocodingApi: %w", err)
	}

	if response.StatusCode() != http.StatusOK {
		return models.LatLong{}, fmt.Errorf("unexpected status code: %d and response: %s", response.StatusCode(), response.String())
	}

	var geoResponse models.GeoResponse
	if err := json.Unmarshal(response.Body(), &geoResponse); err != nil {
		return models.LatLong{}, fmt.Errorf("error decoding response: %w", err)
	}

	if len(geoResponse.Results) == 0 {
		return models.LatLong{}, fmt.Errorf("no results found for city: %s", city)
	}

	return models.LatLong{Lat: geoResponse.Results[0].Latitude, Long: geoResponse.Results[0].Longitude}, nil
}
