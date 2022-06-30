package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mail-app/api/models"
	"mail-app/api/services"
	"net/http"
)

// CreateUser POST /api/user/create/:id (company ID)
func CreateUser(c *gin.Context) {
	user := models.User{}
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

	err = services.GetCompanyByID(&company, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company does not exist"})
		return
	}

	name, email, role := data["name"], data["email"], data["role"]
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	err = services.CreateUser(&user, company, name, email, password, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	c.JSON(http.StatusCreated, "User created successfully")
}

// DeleteUser DELETE /api/user/delete
func DeleteUser(c *gin.Context) {
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	err = services.DeleteUser(&user, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting User"})
		return
	}

	c.JSON(http.StatusOK, "User successfully deleted")
}

// GetUserByID GET /api/user/:id
func GetUserByID(c *gin.Context) {
	user := models.User{}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}

	err = services.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the requested User"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByEmail POST /api/user/email
func GetUserByEmail(c *gin.Context) {
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find the requested User"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserEmail PUT /api/user/update-email/:id
func UpdateUserEmail(c *gin.Context) {
	user := models.User{}
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

	err = services.GetUserByEmail(&user, data["email"])
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	err = services.UpdateUserEmail(&user, id, data["email"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error updating email address"})
		return
	}

	c.JSON(http.StatusOK, "Email updated successfully")
}

// UpdateUserPassword PUT /api/user/update-password
func UpdateUserPassword(c *gin.Context) {
	user := models.User{}
	var data = make(map[string]string)

	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	err = services.GetUserByEmail(&user, data["email"])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}

	email := data["email"]
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	err = services.UpdateUserPassword(&user, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
		return
	}

	c.JSON(http.StatusOK, "Password updated successfully")
}
