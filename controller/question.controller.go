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

func (qc *QuestionController) GetQuestion(ctx *gin.Context) {
	var QuestionID string = ctx.Param("QuestionID")
	user, err := qc.questionService.GetQuestion(QuestionID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (qc *QuestionController) CreateQuestion(ctx *gin.Context) {
	var Question models.Question
	if err := ctx.ShouldBindJSON(&Question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := qc.questionService.CreateQuestion(&Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (qc *QuestionController) UpdateQuestion(ctx *gin.Context) {
	var Question models.Question
	if err := ctx.ShouldBindJSON(&Question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := qc.questionService.UpdateQuestion(&Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (qc *QuestionController) DeleteQuestion(ctx *gin.Context) {
	var QuestionID string = ctx.Param("QuestionID")
	err := qc.questionService.DeleteQuestion(QuestionID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// not implemented
// func (qc *QuestionController) multi_choice(ctx *gin.Context) {

// }

func (qc *QuestionController) RegisterQuestionRouterGroup(rg *gin.RouterGroup) {
	questionrouter := rg.Group("/question")
	questionrouter.GET("/GetQuestion/:id", qc.GetQuestion)
	questionrouter.POST("/CreateQuestion", qc.CreateQuestion)
	questionrouter.PATCH("/UpdateQuestion", qc.UpdateQuestion)
	questionrouter.DELETE("/DeleteQuestion", qc.DeleteQuestion)
	// questionrouter.GET("/multi_choice", qc.multi_choice)
}
