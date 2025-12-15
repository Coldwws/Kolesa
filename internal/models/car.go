package models

import "time"

type Car struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Title        string    `json:"title" db:"title"`
	Brand        string    `json:"brand" db:"brand"`
	Model        string    `json:"model" db:"model"`
	Description  string    `json:"description" db:"description"`
	Price        int64     `json:"price" db:"price"`
	Year         int       `json:"year" db:"year"`
	Mileage      int       `json:"mileage" db:"mileage"`
	EngineType   string    `json:"engine_type" db:"engine_type"`
	EngineVolume float32   `json:"engine_volume" db:"engine_volume"`
	Transmission string    `json:"transmission" db:"transmission"`
	DriveType    string    `json:"drive_type" db:"drive_type"`
	BodyType     string    `json:"body_type" db:"body_type"`
	Color        string    `json:"color" db:"color"`
	Steering     string    `json:"steering" db:"steering"`
	City         string    `json:"city" db:"city"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
