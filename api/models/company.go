package models

import "time"

type Company struct {
	ID        int       `json:"id" gorm:"primaryKey;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Secret    []byte    `json:"secret" gorm:"not null"`
}
