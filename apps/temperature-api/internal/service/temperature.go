package service

import (
	"math/rand"
	"temperature-api/internal/domain"
	"time"
)

type TemperatureServiceInterface interface {
	GetTemperatureByLocation(location string) (*domain.TemperatureData, error)
	GetTemperatureBySensorID(sensorId string) (*domain.TemperatureData, error)
}

var _ TemperatureServiceInterface = (*TemperatureService)(nil)

type TemperatureService struct {
}

func NewTemperatureService() TemperatureServiceInterface {
	return &TemperatureService{}
}

func (s *TemperatureService) GetTemperatureByLocation(location string) (*domain.TemperatureData, error) {

	var sensorId string

	switch location {
	case "Living Room":
		sensorId = "1"
	case "Bedroom":
		sensorId = "2"
	case "Kitchen":
		sensorId = "3"
	default:
		sensorId = "0"
	}

	return generateRandomTemperatureData(location, sensorId), nil
}

func (s *TemperatureService) GetTemperatureBySensorID(sensorId string) (*domain.TemperatureData, error) {

	var location string

	switch sensorId {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}

	return generateRandomTemperatureData(location, sensorId), nil
}

func generateRandomTemperatureData(location string, sensorId string) *domain.TemperatureData {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 15
	max := 35
	randomInt := r.Intn(max-min+1) + min
	randomFloat := float64(randomInt)

	return &domain.TemperatureData{
		Value:       randomFloat,
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "active",
		SensorID:    sensorId,
		SensorType:  "temperature",
		Description: "temperature sensor",
	}
}
