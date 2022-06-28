package models

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Secret    []byte    `json:"secret" gorm:"not null"`
}
