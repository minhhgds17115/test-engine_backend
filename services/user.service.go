package services

import "example.com/m/v2/models"

type userService interface {
	CreateUser(models.Users) error
	getUser(models.Users) (models.Users, error)
	getAllUsers() ([]models.Users, error)
	updateUser(models.Users) error
	deleteUser(models.Users) error
}
