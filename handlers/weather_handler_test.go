package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zoltan-nz/weather-forecast-go/models"
	"github.com/zoltan-nz/weather-forecast-go/services"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWeatherHandler_GetWeather(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := services.NewMockWeatherService()
	logger := log.New(os.Stdout, "weather_hander_test: ", log.LstdFlags)
	handler := NewWeatherHandler(mockService, logger)

	tests := []struct {
		name       string
		cityParam  string
		setupMock  func(*services.MockWeatherService)
		wantStatus int
		wantErr    bool
	}{
		{
			name:      "Valid city returns weather",
			cityParam: "London",
			setupMock: func(m *services.MockWeatherService) {
				m.ShouldError = false
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name:      "Empty city returns error",
			cityParam: "",
			setupMock: func(m *services.MockWeatherService) {
				m.ShouldError = false
			},
			wantStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name:      "Service error returns internal server error",
			cityParam: "London",
			setupMock: func(m *services.MockWeatherService) {
				m.ShouldError = true
			},
			wantStatus: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mock and setup for this test
			mockService.CallCount = 0
			tt.setupMock(mockService)

			// Create test context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "city", Value: tt.cityParam}}

			// Call handler
			handler.GetWeather(c)

			// Assert status code
			if w.Code != tt.wantStatus {
				t.Errorf("GetWeather() status = %v, want %v", w.Code, tt.wantStatus)
			}

			// For success cases, verify response structure
			if tt.wantStatus == http.StatusOK {
				var response models.WeatherData
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}

				if len(response.HourlyTemperatures) == 0 {
					t.Error("Expected non-empty weather data")
				}

				// Verify mock was called exactly once
				if mockService.CallCount != 1 {
					t.Errorf("Expected service to be called once, got %d calls", mockService.CallCount)
				}
			}

			// For error cases, verify error response
			if tt.wantErr {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}

				if _, exists := errorResponse["error"]; !exists {
					t.Error("Expected error message in response")
				}
			}
		})
	}
}
