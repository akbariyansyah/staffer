package employee

import "time"

type Employee struct {
	ID         string    `json:"id"`
	Name       string    `json:"Name"`
	BirthDate  time.Time `json:"birth_date"`
	PositionID string    `json:"position_id"`
	IDNumber   string    `json:"id_number"`
	Gender     string    `json:"gender"`
}
