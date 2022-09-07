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

func (qc *QuestionsController) GetQuestions(ctx *gin.Context) {
	var QuestionsID string = ctx.Param("QuestionsID")
	user, err := qc.QuestionsService.GetQuestions(QuestionsID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (qc *QuestionsController) CreateQuestions(ctx *gin.Context) {
	var Questions models.Questions
	if err := ctx.ShouldBindJSON(&Questions); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := qc.QuestionsService.CreateQuestions(&Questions)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (qc *QuestionsController) UpdateQuestions(ctx *gin.Context) {
	var Questions models.Questions
	if err := ctx.ShouldBindJSON(&Questions); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := qc.QuestionsService.UpdateQuestions(&Questions)
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
// func (qc *QuestionsController) multi_choice(ctx *gin.Context) {

// }

func (qc *QuestionsController) RegisterQuestionsRouterGroup(rg *gin.RouterGroup) {
	Questionsrouter := rg.Group("/Questions")
	// Questionsrouter.GET("",qc.GetAllQuestions)
	Questionsrouter.GET("/:id", qc.GetQuestions)
	Questionsrouter.POST("/", qc.CreateQuestions)
	Questionsrouter.PATCH("/:id", qc.UpdateQuestions)
	Questionsrouter.DELETE("/:id", qc.DeleteQuestions)

}
