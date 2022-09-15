package controller

import (
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	testService *services.TestServiceImpl
}

func NewTestController(testService *services.TestServiceImpl) *TestController {
	return &TestController{
		testService: testService,
	}
}

func (tc *TestController) GetAllTest(ctx *gin.Context) {
	Test, err := tc.testService.GetAllTest()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Test)
}

func (tc *TestController) GetTestID(ctx *gin.Context) {
	testId, _ := strconv.Atoi(ctx.Param("id"))
	user, err := tc.testService.GetTestID(&testId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (tc *TestController) UpdateTest(ctx *gin.Context) {
	var test models.Test
	if err := ctx.ShouldBindJSON(&test); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.testService.UpdateTest(&test)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TestController) StoreAnswer(ctx *gin.Context) {
	var UserAnswer models.UserAnswer
	var Global models.Global
	var ReturnedUserInformation models.ReturnedUserInformation
	var history models.History
	var result models.Result
	var stats models.Stats

	if err := ctx.ShouldBindJSON(&UserAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := tc.testService.StoreAnswer(&Global, &ReturnedUserInformation, &UserAnswer, &history, &result, &stats)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TestController) RegisterTestRouterGroup(rg *gin.RouterGroup) {
	testroute := rg.Group("/test")
	testroute.PATCH(":id", tc.UpdateTest)
	testroute.GET("/", tc.GetAllTest)
	testroute.GET("/:id", tc.GetTestID)
	testroute.POST("/answer", tc.StoreAnswer)
	// testroute.POST("/history", tc.StoreHistory)
}
