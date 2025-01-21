package repositories

import (
	"gouldcs/gouldservice/internal/models"
)

type CityRepository interface {
	GetById(id uint) (*models.City, error)
	Create(city *models.City) error
	Update(city *models.City) error
	Delete(id uint) error
}