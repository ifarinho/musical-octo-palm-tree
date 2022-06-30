package router

import (
	"github.com/gin-gonic/gin"
	"mail-app/api"
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
	u.PUT("/update-email/:id", api.UpdateUserEmail)
	u.PUT("/update-password", api.UpdateUserPassword)
	u.DELETE("/delete", api.DeleteUser)

	/* mail routes */
	m := r.Group("/api/mail")

	m.GET("/user-sent/:id", api.GetSentMails)
	m.POST("/user-received", api.GetReceivedMails)
	m.POST("/send", api.SendEmail)
	m.DELETE("/:id", api.DeleteMail)

	return r
}
