package services

import (
	"gouldcs/gouldservice/internal/models"
	repositories "gouldcs/gouldservice/internal/repositories/city"
)


type CityService struct {
	repo repositories.CityRepository
}

func NewCityService(repo repositories.CityRepository) *CityService {
	return &CityService{repo: repo}
}

func (s *CityService) GetCityById(id uint) (*models.City, error) {
	return s.repo.GetById(id)
}

func (s *CityService) CreateCity(city *models.City) error {
	return s.repo.Create(city)
}

func (s *CityService) UpdateCity(city *models.City) error {
    return s.repo.Update(city)
}

func (s *CityService) DeleteCity(id uint) error {
    return s.repo.Delete(id)
}