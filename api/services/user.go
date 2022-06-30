package services

import (
	"github.com/google/uuid"
	"mail-app/api/models"
	"mail-app/db"
)

func CreateUser(user *models.User, company models.Company, name string, email string, pass []byte, role string) error {
	user.ID = uuid.New()
	user.Name = name
	user.Email = email
	user.Password = pass
	user.Company = &company
	user.Role = role

	res := db.DB().Create(user)

	return res.Error
}

func DeleteUser(user *models.User, email string) error {
	res := db.DB().Where("email = ?", email).Delete(&user)

	return res.Error
}

/* get functions */

func GetUserByID(user *models.User, id uuid.UUID) error {
	res := db.DB().Preload("Company").Where("id = ?", id).First(&user)

	return res.Error
}

func GetUserByEmail(user *models.User, email string) error {
	res := db.DB().Preload("Company").Where("email = ?", email).First(&user)

	return res.Error
}

/* update functions */

func UpdateUserEmail(user *models.User, id uuid.UUID, email string) error {
	res := db.DB().Model(&user).Where("id = ?", id).Update("email", email)

	return res.Error
}

func UpdateUserPassword(user *models.User, email string, password []byte) error {
	res := db.DB().Model(&user).Where("email = ?", email).Update("password", password)

	return res.Error
}
