package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/m/v2/models"
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

func (uc *QuestionController) GetQuestion(ctx *gin.Context) {
	var QuestionID string = ctx.Param("QuestionID")
	user, err := uc.questionService.GetQuestion(QuestionID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *QuestionController) CreateQuestion(ctx *gin.Context) {
	var Question models.Question
	if err := ctx.ShouldBindJSON(&Question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.questionService.CreateQuestion(&Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *QuestionController) UpdateQuestion(ctx *gin.Context) {
	var Question models.Question
	if err := ctx.ShouldBindJSON(&Question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.questionService.UpdateQuestion(&Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *QuestionController) DeleteQuestion(ctx *gin.Context) {
	var Question string = ctx.Param("Question")
	err := uc.questionService.DeleteQuestion(Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *QuestionController) multi_choice(ctx *gin.Context) {

}
