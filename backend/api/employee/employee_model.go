package employee

import "time"

type Employee struct {
	ID         string    `json:"id"`
	FullName   string    `json:"full_name" validate:"required"`
	BirthDate  time.Time `json:"birth_date" validate:"required"`
	PositionID string    `json:"position_id" validate:"required"`
	IDNumber   string    `json:"id_number" validate:"required"`
	Gender     string    `json:"gender" validate:"required"`
	IsDelete   string    `json:"is_delete"`
}
