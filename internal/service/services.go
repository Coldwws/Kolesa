package service

import "github.com/Coldwws/kolesa/internal/models"

type CarService interface {
	GetAllCars() ([]models.Car, error)
	GetCarByID(id int64) (*models.Car, error)
	CreateCar(car models.Car) error
	UpdateCar(id int64, updateCar models.UpdateCar) error
	DeleteCar(id int64) (int, error)
}