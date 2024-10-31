package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoltan-nz/weather-forecast-go/services"
	"log"
	"net/http"
)

type WeatherHandler struct {
	weatherService services.WeatherService
	logger         *log.Logger
}

func NewWeatherHandler(ws services.WeatherService, logger *log.Logger) *WeatherHandler {
	return &WeatherHandler{
		weatherService: ws,
		logger:         logger,
	}
}

// GetWeather handles GET /api/weather/:city
func (h *WeatherHandler) GetWeather(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		h.logger.Printf("Empy city parameter received")
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}

	h.logger.Printf("Fetching weather for city %s", city)

	// Get coordinates
	latLong, err := services.FetchLatLong(city)
	if err != nil {
		h.logger.Printf("Error fetching coordinates for city %s: %v", city, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
		return
	}

	// Get weather data
	weather, err := h.weatherService.FetchWeather(latLong)
	if err != nil {
		h.logger.Printf("Error fetching weather for city %s: %v", city, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching weather data"})
		return
	}

	c.JSON(http.StatusOK, weather)
}
