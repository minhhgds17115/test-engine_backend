package controller

import (
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

type AnswerController struct {
	answerService *services.AnswerServiceImpl
}

func NewAnswerController(answerService *services.AnswerServiceImpl) *AnswerController {
	return &AnswerController{
		answerService: answerService,
	}
}

func (ac *AnswerController) GetAnswer(ctx *gin.Context) {
	AnswerId, err := strconv.Atoi(ctx.Param("AnswerId"))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	user, err := ac.answerService.GetAnswer(AnswerId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (ac *AnswerController) DeleteAnswer(ctx *gin.Context) {
	var Answer models.Answer
	err := ac.answerService.DeleteAnswer(&Answer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AnswerController) CreateAnswer(ctx *gin.Context) {
	var Answer models.Answer
	if err := ctx.ShouldBindJSON(&Answer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ac.answerService.CreateAnswer(&Answer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AnswerController) UpdateAnswer(ctx *gin.Context) {
	var Answer models.Answer
	if err := ctx.ShouldBindJSON(&Answer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ac.answerService.UpdateAnswer(&Answer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AnswerController) RegisterAnswerRouterGroup(rg *gin.RouterGroup) {
	answerrouter := rg.Group("/answer")
	// answerrouter.GET("", ac.GetAllAnswer)
	answerrouter.GET("/:id", ac.GetAnswer)
	answerrouter.POST("", ac.CreateAnswer)
	answerrouter.PATCH("/:id", ac.UpdateAnswer)
	answerrouter.DELETE(":id", ac.DeleteAnswer)
	// answerrouter.POST("", ac.PostAnswer)
}
