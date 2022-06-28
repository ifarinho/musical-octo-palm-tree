package models

import "time"

type Mail struct {
	ID         int       `json:"id" gorm:"primaryKey;not null"`
	CreatedAt  time.Time `json:"created_at"`
	FKSender   int       `json:"fk_sender" gorm:"not null"`
	FKReceiver int       `json:"fk_receiver" gorm:"not null"`
	Sender     User      `json:"sender" gorm:"foreignKey:FKSender;not null"`
	Receiver   User      `json:"receiver" gorm:"foreignKey:FKReceiver;not null"`
	Subject    string    `json:"subject" gorm:"not null"`
	Body       string    `json:"body" gorm:"not null"`
}
