package services

import "example.com/m/v2/models"

type UserService interface {
	CreateUser(models.Users) error
	GetUserEmail(models.Users) (models.Users, error)
	GetAllUsers() ([]models.Users, error)
	updateUser(models.Users) error
	deleteUser(models.Users) error
}
