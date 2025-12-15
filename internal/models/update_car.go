package models

type UpdateCar struct {
	Title       *string  `json:"title"`
	Brand       *string  `json:"brand"`
	Model       *string  `json:"model"`
	Description *string  `json:"description"`

	Price    *int64 `json:"price"`
	Year     *int   `json:"year"`
	Mileage  *int   `json:"mileage"`

	EngineType   *string  `json:"engine_type"`
	EngineVolume *float32 `json:"engine_volume"`

	Transmission *string `json:"transmission"`
	DriveType    *string `json:"drive_type"`

	BodyType *string `json:"body_type"`
	Color    *string `json:"color"`

	Steering *string `json:"steering"`

	City   *string `json:"city"`
	Status *string `json:"status"`
}
