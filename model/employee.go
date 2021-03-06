package model

import "time"

type Employee struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	IsMarried bool      `json:"is_married"`
	BirthDate time.Time `json:"birth_date"`
}
