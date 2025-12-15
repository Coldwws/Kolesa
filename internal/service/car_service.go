package service

import (
	"github.com/Coldwws/kolesa/internal/models"
	"github.com/Coldwws/kolesa/internal/repository"
)

type carService struct {
    repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
    return &carService{repo: repo}
}


func (s *carService) GetAllCars() ([]models.Car, error) {
		return s.repo.GetAllCars()
}
func (s *carService) GetCarByID(id int64) (*models.Car, error) {
		return s.repo.GetCarByID(id)
}
func (s *carService) CreateCar(car models.Car) error {
		return s.repo.CreateCar(car)
}
func (s *carService) UpdateCar(id int64, updateCar models.UpdateCar) error {
		return s.repo.UpdateCar(id, updateCar)
}
func (s *carService) DeleteCar(id int64) (int, error) {
		return s.repo.DeleteCar(id)
}
