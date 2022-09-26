package services

import "example.com/m/v2/models"

// Candidate services Interface

type UserService interface {
	CreateUser(models.Candidate) error
	GetUserEmail(models.Candidate) (models.Candidate, error)
	GetAllUsers() ([]models.Candidate, error)
	updateUser(models.Candidate) error
	deleteUser(models.Candidate) error
	GetUserTestID(models.CandidateInformation) error
}
