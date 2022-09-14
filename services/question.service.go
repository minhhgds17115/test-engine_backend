package services

import "example.com/m/v2/models"

type QuestionsService interface {
	CreateQuestions(models.Question) error
	GetQuestions(models.Question) (models.Question, error)
	getAllQuestions() ([]models.Question, error)
	updateQuestions(models.Question) error
	deleteQuestions(models.Question) error
	multi_choice(models.Question) error
	GetTopics(models.Question)
}
