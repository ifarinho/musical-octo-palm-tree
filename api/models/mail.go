package models

import (
	"github.com/google/uuid"
	"time"
)

type Mail struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;not null"`
	CreatedAt time.Time `json:"created_at"`
	FKSender  uuid.UUID `json:"fk_sender" gorm:"not null"`
	Sender    *User     `json:"sender" gorm:"foreignKey:FKSender;not null"`
	Receiver  string    `json:"receiver" gorm:"not null"`
	Subject   string    `json:"subject" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
}
