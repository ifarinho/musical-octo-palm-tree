package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"io"
	"mail-app/api/models"
	"mail-app/config"
)

func creds() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, ""),
	})

	return sess
}

func UploadFileToBucket(file io.ReadSeeker) error {
	_, err := s3.New(creds()).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(uuid.NewString()),
		Body:   file,
	})

	return err
}

func SendEmailSES(mail *models.Mail, receivers []string, attachment io.ReadSeeker) error {
	svc := ses.New(creds())
	e := email.NewEmail()

	e.From = mail.Sender.Email
	e.To = receivers
	e.Cc = receivers
	e.Subject = mail.Subject
	e.Text = []byte(mail.Body)

	if attachment != nil {
		_, err := e.Attach(attachment, "attachment.pdf", "")
		return err
	}

	bytes, err := e.Bytes()

	input := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: bytes,
		},
		Destinations: aws.StringSlice(receivers),
		Source:       aws.String(mail.Sender.Email),
	}

	_, err = svc.SendRawEmail(input)

	return err
}
