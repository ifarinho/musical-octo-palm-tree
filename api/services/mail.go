package services

import (
	"github.com/google/uuid"
	"mail-app/api/models"
	"mail-app/db"
)

/* create and delete */

func CreateMail(mail *models.Mail, sender models.User, receiver string, subject string, body string) error {
	mail.ID = uuid.New()
	mail.Sender = &sender
	mail.Receiver = receiver
	mail.Subject = subject
	mail.Body = body

	res := db.DB().Create(mail)

	return res.Error
}

func DeleteMail(mail *models.Mail, id uuid.UUID) error {
	res := db.DB().Where("id = ?", id).Delete(&mail)

	return res.Error
}

/* get functions */

func GetSentMails(mails *[]models.Mail, id uuid.UUID) error {
	res := db.DB().Preload("Sender.Company").Where("fk_sender = ?", id).Find(&mails)

	return res.Error
}

func GetReceivedMails(mails *[]models.Mail, email string) error {
	res := db.DB().Preload("Sender.Company").Where("receiver = ?", email).Find(&mails)

	return res.Error
}
