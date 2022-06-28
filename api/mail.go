package api

import (
	"electro3-project-go/api/models"
	"electro3-project-go/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

/* send and delete */

func SendEmail(c *gin.Context) {
	mail := models.Mail{}
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email data"})
		return
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user email"})
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	err = services.CreateMail(&mail, user, data["receiver"], data["subject"], data["body"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating email"})
		return
	}

	err = services.SendMail(&mail, data["password"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		return
	}

	c.JSON(http.StatusOK, mail)
}

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
	}

	c.JSON(http.StatusOK, "Email deleted successfully")
}

/* get functions */

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
