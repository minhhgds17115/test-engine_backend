package controller

import (
	"github.com/gin-gonic/gin"

	"example.com/m/v2/services"
)

type QuestionController struct {
	questionService *services.QuestionServiceImpl
}

func NewQuestionController(questionService *services.QuestionServiceImpl) *QuestionController {
	return &QuestionController{
		questionService: questionService,
	}
}

func (uc *UserController) GetQuestion(ctx *gin.Context) {

}

func (uc *UserController) CreateQuestion(ctx *gin.Context) {

}

func (uc *UserController) updateQuestion(ctx *gin.Context) {

}

func (uc *UserController) DeleteQuestion(ctx *gin.Context) {

}

func (uc *UserController) multi_choice(ctx *gin.Context) {

}

func (uc *UserController) GetTopic(ctx *gin.Context) {
}
