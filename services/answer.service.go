package services

import "example.com/m/v2/models"

// Answer services Interface
type AnswerService interface {
	getAnswer(models.Answer) (*models.Answer, error)
	createAnswer(models.Answer) (*models.Answer, error)
	updateAnswer(models.Answer) (*models.Answer, error)
	deleteAnswer(models.Answer) error
	postAnswer(models.Answer) error
}
