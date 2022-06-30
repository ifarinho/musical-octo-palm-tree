package services

import (
	"github.com/google/uuid"
	"mail-app/api/models"
	"mail-app/db"
)

/* create and delete */

func CreateCompany(company *models.Company, name string, email string, secret []byte) error {
	company.ID = uuid.New()
	company.Name = name
	company.Email = email
	company.Secret = secret

	res := db.DB().Create(company)

	return res.Error
}

func DeleteCompany(company *models.Company, email string) error {
	res := db.DB().Where("email = ?", email).Delete(&company)

	return res.Error
}

/* get functions */

func GetCompanyByID(company *models.Company, id uuid.UUID) error {
	res := db.DB().Where("id = ?", id).First(&company)

	return res.Error
}

func GetCompanyByEmail(company *models.Company, email string) error {
	res := db.DB().Where("email = ?", email).First(&company)

	return res.Error
}

/* update functions */

func UpdateCompanyEmail(company *models.Company, id uuid.UUID, email string) error {
	res := db.DB().Model(&company).Where("id = ?", id).Update("email", email)

	return res.Error
}

func UpdateCompanySecretPhrase(company *models.Company, email string, secret []byte) error {
	res := db.DB().Model(&company).Where("email = ?", email).Update("secret", secret)

	return res.Error
}
