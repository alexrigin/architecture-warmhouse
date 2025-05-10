package handler

import (
	"temperature-api/internal/service"
)

type TemperatureSensorHandler struct {
	srv service.TemperatureServiceInterface
}
