package router

import (
	"electro3-project-go/api"
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	r := gin.Default()

	/* company routes */
	c := r.Group("/api/company")

	c.GET("/:id", api.GetCompanyByID)
	c.POST("/email", api.GetCompanyByEmail)
	c.POST("/create", api.CreateCompany)
	c.PUT("/update-email/:id", api.UpdateCompanyEmail)
	c.PUT("/update-secret-phrase", api.UpdateCompanySecretPhrase)
	c.DELETE("/delete", api.DeleteCompany)

	/* user routes */
	u := r.Group("/api/user")

	u.GET("/:id", api.GetUserByID)
	u.POST("/email", api.GetUserByEmail)
	u.POST("/create/:id", api.CreateUser)
	c.PUT("/update-email/:id", api.UpdateUserEmail)
	c.PUT("/update-password", api.UpdateUserPassword)
	u.DELETE("/delete", api.DeleteUser)

	return r
}
