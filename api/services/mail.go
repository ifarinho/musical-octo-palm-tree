package services

import (
	"electro3-project-go/api/models"
	"electro3-project-go/db"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

/* create and send */

func CreateMail(mail *models.Mail, sender models.User, receiver string, subject string, body string) error {
	mail.ID = uuid.New()
	mail.Sender = sender
	mail.Receiver = receiver
	mail.Subject = subject
	mail.Body = body

	res := db.DB().Create(mail)

	return res.Error
}

func SendMail(mail *models.Mail, password string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.Sender.Email)
	m.SetHeader("To", mail.Receiver)
	m.SetAddressHeader("Cc", mail.Sender.Email, mail.Sender.Name)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	d := gomail.NewDialer("smtp.example.com", 587, mail.Sender.Email, password)
	err := d.DialAndSend(m)

	return err
}

/* delete mails */

func DeleteMail(mail *models.Mail, id uuid.UUID) error {
	res := db.DB().Where("id = ?", id).Delete(&mail)

	return res.Error
}

/* get functions */

func GetSentMails(mails *[]models.Mail, id uuid.UUID) error {
	res := db.DB().Where("fk_sender = ?", id).Find(&mails)

	return res.Error
}

func GetReceivedMails(mails *[]models.Mail, email string) error {
	res := db.DB().Where("receiver = ?", email).Find(&mails)

	return res.Error
}
