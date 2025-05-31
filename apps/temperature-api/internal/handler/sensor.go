package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temperature-api/internal/service"
)

type SensorHandler struct {
	srv service.TemperatureServiceInterface
}

// NewSensorHandler creates a new SensorHandler
func NewSensorHandler(temperatureService service.TemperatureServiceInterface) *SensorHandler {
	return &SensorHandler{
		srv: temperatureService,
	}
}

// RegisterRoutes registers the sensor routes
func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
	sensors := router.Group("/temperature")
	{
		sensors.GET("/", h.GetTemperatureByLocation)
		sensors.GET("/:id", h.GetTemperatureBySensorID)
	}
}

// GetTemperatureByLocation handles GET /api/v1/temperature/:location
func (h *SensorHandler) GetTemperatureByLocation(c *gin.Context) {
	location := c.Query("location")

	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location required"})
		return
	}

	tempData, err := h.srv.GetTemperatureByLocation(location)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unexpected error"})
		return
	}

	c.JSON(http.StatusOK, tempData)
}

// GetTemperatureBySensorID handles GET /api/v1/temperature/:id
func (h *SensorHandler) GetTemperatureBySensorID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	tempData, err := h.srv.GetTemperatureBySensorID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unexpected error"})
		return
	}

	c.JSON(http.StatusOK, tempData)
}
