package models

import "time"

type Base struct {
	ID          string    `json:"id"yaml:"id"`
	Name        string    `json:"name"yaml:"name"`
	Description string    `json:"description"yaml:"description"`
	CreatedAt   time.Time `json:"created_at"yaml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"yaml:"updated_at"`
}
