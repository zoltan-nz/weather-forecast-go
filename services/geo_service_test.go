package services

import (
	"github.com/zoltan-nz/weather-forecast-go/models"
	"testing"
)

// FetchLatLong is a function that fetches the latitude and longitude of a city
// using the OpenMeteoGeocoding API.
// It should return the latitude and longitude objects (LatLong) if the city is found.
func TestFetchLatLong(t *testing.T) {
	tests := []struct {
		name    string
		city    string
		want    models.LatLong
		wantErr bool
	}{
		{
			name:    "Valid city",
			city:    "London",
			want:    models.LatLong{Lat: 51.50853, Long: -0.12574},
			wantErr: false,
		},
		{
			name:    "Invalid city",
			city:    "NonExistentCity",
			want:    models.LatLong{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			{
			}
			latLong, err := FetchLatLong(tt.city)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchLatLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if latLong.Lat != tt.want.Lat {
				t.Errorf("FetchLatLong() lat = %v, wantLat %v", latLong.Lat, tt.want.Lat)
			}
			if latLong.Long != tt.want.Long {
				t.Errorf("FetchLatLong() long = %v, wantLong %v", latLong.Long, tt.want.Long)
			}
		})
	}
}
