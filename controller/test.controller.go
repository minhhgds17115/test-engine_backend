package controller

import (
	"net/http"
	"strconv"

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
	testId, _ := strconv.Atoi(ctx.Param("test_id"))
	user, err := tc.testService.GetTestID(&testId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (tc *TestController) RegisterTestRouterGroup(rg *gin.RouterGroup) {
	testroute := rg.Group("/test")

	testroute.GET("/getAllTest", tc.GetAllTest)
	testroute.GET("/GetTestID/:id", tc.GetTestID)
}
