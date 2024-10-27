package models

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
