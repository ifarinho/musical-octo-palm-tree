package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  []byte    `json:"password" gorm:"not null"`
	FKCompany int       `json:"fk_company" gorm:"not null"`
	Company   Company   `json:"company" gorm:"foreignKey:FKCompany;not null"`
	Role      string    `json:"role" gorm:"not null"`
}
