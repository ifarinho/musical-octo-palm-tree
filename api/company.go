package api

import (
	"electro3-project-go/api/models"
	"electro3-project-go/api/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

/* create and delete */

func CreateCompany(c *gin.Context) {
	company := models.Company{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	name, email := data["name"], data["email"]
	secret, _ := bcrypt.GenerateFromPassword([]byte(data["secret"]), 14)

	err = services.GetCompanyByEmail(&company, email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is already a company registered with this email"})
		return
	}

	err = services.CreateCompany(&company, name, email, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Company creation failed"})
		return
	}

	c.JSON(http.StatusCreated, "Company successfully created")
}

func DeleteCompany(c *gin.Context) {
	company := models.Company{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetCompanyByEmail(&company, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company does not exist"})
		return
	}

	err = bcrypt.CompareHashAndPassword(company.Secret, []byte(data["secret"]))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect secret phrase"})
		return
	}

	err = services.DeleteCompany(&company, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting Company"})
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted Company")
}

/* get functions */

func GetCompanyByID(c *gin.Context) {
	company := models.Company{}

	id, _ := uuid.Parse(c.Param("id"))

	err := services.GetCompanyByID(&company, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the requested company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    company.ID,
		"name":  company.Name,
		"email": company.Email,
	})
}

func GetCompanyByEmail(c *gin.Context) {
	company := models.Company{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetCompanyByEmail(&company, data["email"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the requested company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    company.ID,
		"name":  company.Name,
		"email": company.Email,
	})
}

/* update functions */

func UpdateCompanyEmail(c *gin.Context) {
	company := models.Company{}
	var data = make(map[string]string)

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}

	err = c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetCompanyByEmail(&company, data["email"])
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	err = services.UpdateCompanyEmail(&company, id, data["email"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error updating email address"})
		return
	}

	c.JSON(http.StatusOK, "Email updated successfully")
}

func UpdateCompanySecretPhrase(c *gin.Context) {
	company := models.Company{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetCompanyByEmail(&company, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company does not exist"})
		return
	}

	email := data["email"]
	secret, _ := bcrypt.GenerateFromPassword([]byte(data["secret"]), 14)

	err = services.UpdateCompanySecretPhrase(&company, email, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating secret phrase"})
		return
	}

	c.JSON(http.StatusOK, "Secret phrase updated successfully")
}
