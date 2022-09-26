package services

import "example.com/m/v2/models"

// Test services Interface

type TestService interface {
	GetAllTest([]models.Test, error)
	CreateTest(models.Test) error
	GetTestID(models.Test) error
	UpdateTest(models.Test) error
	DeleteTest(models.Test) error
	StoreUserInfo(models.CandidateInformation) error
	ReturnAnswer(models.ReturnedAnswer) error
	StoreTestCandidate(models.Test) error
}
