package services

import "example.com/m/v2/models"

type TestService interface {
	GetAllTest([]models.Test, error)
	CreateTest(models.Test) error
	GetTestID(models.Test) error
	UpdateTest(models.Test) error
	DeleteTest(models.Test) error
	StoreAnswer([]models.ReturnedAnswer) error
	StoreHistory(models.UserAnswer) error
}
