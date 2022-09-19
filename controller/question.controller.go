package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
)

type QuestionsController struct {
	QuestionsService *services.QuestionsServiceImpl
}

func NewQuestionsController(QuestionsService *services.QuestionsServiceImpl) *QuestionsController {
	return &QuestionsController{
		QuestionsService: QuestionsService,
	}
}

func (qc *QuestionsController) GetQuestionsByID(ctx *gin.Context) {
	var QuestionsID string = ctx.Param("QuestionsID")
	user, err := qc.QuestionsService.GetQuestionsByID(QuestionsID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (qc *QuestionsController) CreateQuestions(ctx *gin.Context) {
	var Question models.Questions
	if err := ctx.ShouldBindJSON(&Question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := qc.QuestionsService.CreateQuestions(&Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (qc *QuestionsController) UpdateQuestions(ctx *gin.Context) {
	var Question models.Questions
	if err := ctx.ShouldBindJSON(&Question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := qc.QuestionsService.UpdateQuestions(&Question)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (qc *QuestionsController) DeleteQuestions(ctx *gin.Context) {
	var QuestionsID string = ctx.Param("id")
	err := qc.QuestionsService.DeleteQuestions(QuestionsID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// not implemented
// func (qc *QuestionsController) GetAllQuestions(ctx *gin.Context) {
// 	Test, err := qc.QuestionsService.GetAllQuestions()
// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Test)
// }

func (qc *QuestionsController) RegisterQuestionsRouterGroup(rg *gin.RouterGroup) {
	Questionsrouter := rg.Group("/Question")
	// Questionsrouter.GET("",qc.GetAllQuestions)
	Questionsrouter.GET("/:id", qc.GetQuestionsByID)
	Questionsrouter.POST("/", qc.CreateQuestions)
	Questionsrouter.PATCH("/:id", qc.UpdateQuestions)
	Questionsrouter.DELETE("/:id", qc.DeleteQuestions)

}
