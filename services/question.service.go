package services

import "example.com/m/v2/models"

type questionService interface {
	CreateQuestion(models.Question) error
	GetQuestion(models.Question) (models.Question, error)
	getAllQuestion() ([]models.Question, error)
	updateQuestion(models.Question) error
	deleteQuestion(models.Question) error
	multi_choice(models.Question) error
	GetTopics(models.Question)
}
