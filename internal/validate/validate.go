package validate

import (
		"time"
		"github.com/Coldwws/kolesa/internal/models"
)

func ValidateCar(car models.Car) (bool, string) {
    currentYear := time.Now().Year()
    if car.Year < 1950 || car.Year > currentYear+1 {
        return false, "invalid year"
    }
    if car.Price <= 0 {
        return false, "price must be greater than 0"
    }
    if car.Mileage < 0 {
        return false, "mileage must be >= 0"
    }
    if car.Brand == "" || car.Model == "" {
        return false, "brand and model are required"
    }
    return true, ""
}

func ValidateUpdateCar(car models.UpdateCar) (bool, string) {
	currentYear := time.Now().Year()

	if car.Year != nil {
		if *car.Year < 1950 || *car.Year > currentYear+1 {
			return false, "invalid year"
		}
	}

	if car.Price != nil && *car.Price <= 0 {
		return false, "price must be greater than 0"
	}

	if car.Mileage != nil && *car.Mileage < 0 {
		return false, "mileage must be >= 0"
	}

	if car.Brand != nil && *car.Brand == "" {
		return false, "brand cannot be empty"
	}

	if car.Model != nil && *car.Model == "" {
		return false, "model cannot be empty"
	}

	return true, ""
}

