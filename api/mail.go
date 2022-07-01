package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mail-app/api/models"
	"mail-app/api/services"
	"mime/multipart"
	"net/http"
)

// SendEmail POST /api/mail/send
func SendEmail(c *gin.Context) {
	var file multipart.File

	mail := models.Mail{}
	user := models.User{}
	data := &struct {
		Email      string                `form:"email"`
		Password   string                `form:"password"`
		Receivers  []string              `form:"receivers"`
		Subject    string                `form:"subject"`
		Body       string                `form:"body"`
		Attachment *multipart.FileHeader `form:"attachment"`
	}{}

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email data"})
		return
	}

	err = services.GetUserByEmail(&user, data.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sender email"})
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	for _, r := range data.Receivers {
		err = services.CreateMail(&mail, user, r, data.Subject, data.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating email"})
			return
		}
	}

	if data.Attachment != nil {
		file, err = data.Attachment.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading file"})
			return
		}

		defer file.Close()

		err = services.UploadFileToBucket(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
			return
		}
	}

	err = services.SendEmailSES(&mail, data.Receivers, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		return
	}

	c.JSON(http.StatusOK, "Email successfully sent")
}

// DeleteMail DELETE /api/mail/:id
func DeleteMail(c *gin.Context) {
	mail := models.Mail{}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}

	err = services.DeleteMail(&mail, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting mail"})
		return
	}

	c.JSON(http.StatusOK, "Email deleted successfully")
}

// GetSentMails GET /api/mail/user-sent/:id
func GetSentMails(c *gin.Context) {
	user := models.User{}
	var mails []models.Mail

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}

	err = services.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	err = services.GetSentMails(&mails, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting mails"})
		return
	}

	c.JSON(http.StatusOK, mails)
}

// GetReceivedMails POST /api/mail/user-received
func GetReceivedMails(c *gin.Context) {
	var mails []models.Mail
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email data"})
		return
	}

	err = services.GetReceivedMails(&mails, data["email"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting mails"})
		return
	}

	c.JSON(http.StatusOK, mails)
}
