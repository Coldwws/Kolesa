package repository

import (
	"strconv"
	"strings"

	"github.com/Coldwws/kolesa/internal/models"
	"github.com/jmoiron/sqlx"
)

type CarPostgres struct {
	db *sqlx.DB
}

func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}

}

func (r *CarPostgres) GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	query := `SELECT id, user_id, title, brand, model, description, price, year, mileage, engine_type, engine_volume, transmission, drive_type, body_type, color, steering, city, status, created_at, updated_at FROM cars`
	err := r.db.Select(&cars, query)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *CarPostgres) GetCarByID(id int64) (*models.Car, error) {
	var car models.Car
	query := `SELECT id, user_id, title, brand, model, description, price, year, mileage, engine_type, engine_volume, transmission, drive_type, body_type, color, steering, city, status, created_at, updated_at FROM cars WHERE id = $1`
	err := r.db.Get(&car, query, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *CarPostgres) CreateCar(car models.Car) error {
	var id int64
	query := `INSERT INTO cars (user_id, title, brand, model, description, price, year, mileage, engine_type, engine_volume, transmission, drive_type, body_type, color, steering, city, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, NOW(), NOW()) RETURNING id`
	err := r.db.QueryRow(query, car.UserID, car.Title, car.Brand, car.Model, car.Description, car.Price, car.Year, car.Mileage, car.EngineType, car.EngineVolume, car.Transmission, car.DriveType, car.BodyType, car.Color, car.Steering, car.City, car.Status).Scan(&id)
	if err != nil {
		return err
	}
	car.ID = id
	return nil

}

func (r *CarPostgres) UpdateCar(id int64, updateCar models.UpdateCar) error {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if updateCar.Title != nil {
		setParts = append(setParts, "title = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Title)
		argIndex++
	}
	if updateCar.Brand != nil {
		setParts = append(setParts, "brand = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Brand)
		argIndex++
	}
	if updateCar.Model != nil {
		setParts = append(setParts, "model = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Model)
		argIndex++
	}
	if updateCar.Description != nil {
		setParts = append(setParts, "description = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Description)
		argIndex++
	}
	if updateCar.Price != nil {
		setParts = append(setParts, "price = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Price)
		argIndex++
	}
	if updateCar.Year != nil {
		setParts = append(setParts, "year = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Year)
		argIndex++
	}
	if updateCar.Mileage != nil {
		setParts = append(setParts, "mileage = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Mileage)
		argIndex++
	}
	if updateCar.EngineType != nil {
		setParts = append(setParts, "engine_type = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.EngineType)
		argIndex++
	}
	if updateCar.EngineVolume != nil {
		setParts = append(setParts, "engine_volume = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.EngineVolume)
		argIndex++
	}
	if updateCar.Transmission != nil {
		setParts = append(setParts, "transmission = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Transmission)
		argIndex++
	}
	if updateCar.DriveType != nil {
		setParts = append(setParts, "drive_type = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.DriveType)
		argIndex++
	}
	if updateCar.BodyType != nil {
		setParts = append(setParts, "body_type = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.BodyType)
		argIndex++
	}
	if updateCar.Color != nil {
		setParts = append(setParts, "color = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Color)
		argIndex++
	}
	if updateCar.Steering != nil {
		setParts = append(setParts, "steering = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Steering)
		argIndex++
	}
	if updateCar.City != nil {
		setParts = append(setParts, "city = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.City)
		argIndex++
	}
	if updateCar.Status != nil {
		setParts = append(setParts, "status = $"+strconv.Itoa(argIndex))
		args = append(args, *updateCar.Status)
		argIndex++
	}

	if len(setParts) == 0 {
		return nil
	}

	setParts = append(setParts, "updated_at = NOW()")
	query := "UPDATE cars SET " + strings.Join(setParts, ", ") + " WHERE id = $" + strconv.Itoa(argIndex)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
func (r *CarPostgres) DeleteCar(id int64) (int, error) {
	query := `DELETE FROM cars WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
