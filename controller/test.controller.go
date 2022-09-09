package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	var Results models.Results

	if err := ctx.ShouldBindJSON(&Results); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.testService.StoreAnswer(&Results)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TestController) StoreHistory(ctx *gin.Context) {
	var History models.History
	timestamp := time.Now()

	if err := ctx.ShouldBindJSON(&History); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.testService.StoreHistory(&History)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	fmt.Println("time completed successfully quiz :", timestamp)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (tc *TestController) RegisterTestRouterGroup(rg *gin.RouterGroup) {
	testroute := rg.Group("/test")
	testroute.PATCH(":id", tc.UpdateTest)
	testroute.GET("/", tc.GetAllTest)
	testroute.GET("/:id", tc.GetTestID)
	testroute.POST("/answer", tc.StoreAnswer)
	testroute.POST("/history", tc.StoreHistory)
}
