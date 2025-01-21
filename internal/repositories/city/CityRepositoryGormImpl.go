package repositories

import (
	"gouldcs/gouldservice/internal/models"

	"gorm.io/gorm"
)

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{db: db}
}

func (r *cityRepository) GetById(id uint) (*models.City, error) {
	var city models.City
	if err := r.db.First(&city, id).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (r *cityRepository) Create(city *models.City) error {
	return r.db.Create(city).Error
}

func (r *cityRepository) Update(city *models.City) error {
    return r.db.Save(city).Error
}

func (r *cityRepository) Delete(id uint) error {
    return r.db.Delete(&models.City{}, id).Error
}