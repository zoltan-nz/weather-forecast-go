package api

import (
	"testing"
)

func TestFetchLatLong(t *testing.T) {
	tests := []struct {
		name     string
		city     string
		wantLat  float64
		wantLong float64
		wantErr  bool
	}{
		{
			name:     "Valid city",
			city:     "London",
			wantLat:  51.50853,
			wantLong: -0.12574,
			wantErr:  false,
		},
		{
			name:     "Invalid city",
			city:     "NonExistentCity",
			wantLat:  0,
			wantLong: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lat, long, err := FetchLatLong(tt.city)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchLatLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if lat != tt.wantLat {
				t.Errorf("FetchLatLong() lat = %v, wantLat %v", lat, tt.wantLat)
			}
			if long != tt.wantLong {
				t.Errorf("FetchLatLong() long = %v, wantLong %v", long, tt.wantLong)
			}
		})
	}
}
