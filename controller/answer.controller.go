package controller

import (
	"net/http"

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

func (uc *AnswerController) getAnswer(ctx *gin.Context) {
	var Answerid string = ctx.Param("Answerid")
	user, err := uc.answerService.getAnswer(Answerid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
