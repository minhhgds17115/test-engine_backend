package services

import "example.com/m/v2/models"

type testService interface {
	GetAllTest([]models.Test, error)
	CreateTest(models.Test) error
}
