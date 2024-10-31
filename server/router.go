package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zoltan-nz/weather-forecast-go/handlers"
	"github.com/zoltan-nz/weather-forecast-go/services"
	"log"
)

type Server struct {
	router *gin.Engine
	logger *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		router: gin.Default(),
		logger: logger,
	}
}

func (s *Server) SetupRoutes() {
	weatherService := services.NewWeatherService(s.logger)
	weatherHandler := handlers.NewWeatherHandler(weatherService, s.logger)

	api := s.router.Group("/api")
	{
		api.GET("/weather/:city", weatherHandler.GetWeather)
	}

	// Add a new route for the health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Add root route
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Weather Forecast API"})
	})
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
