package services

import "example.com/m/v2/models"

type QuestionsService interface {
	CreateQuestions(models.Questions) error
	GetQuestions(models.Questions) (models.Questions, error)
	getAllQuestions() ([]models.Questions, error)
	updateQuestions(models.Questions) error
	deleteQuestions(models.Questions) error
	multi_choice(models.Questions) error
	GetTopics(models.Questions)
}
