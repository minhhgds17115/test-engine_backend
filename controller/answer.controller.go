package controller

import (
	"example.com/m/v2/services"
)

type AnswerController struct {
	questionService *services.AnswerServiceImpl
}

func NewAnswerController(answerService *services.AnswerServiceImpl) *AnswerController {
	return &AnswerController{
		// answerService: answerService,
	}
}
