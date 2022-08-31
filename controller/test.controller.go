package controller

import (
	"net/http"

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

func (tc *TestController) CreateTest(ctx *gin.Context) {
	// var TestID models.Test
	// if err := ctx.ShouldBindJSON(&TestID); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// err := tc.testService.CreateTest(&TestID)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{"message": "success"})
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
	// var testId string = ctx.Param("testId")
	// user, err := tc.testService.GetTestID(&testId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, user)
}

func (tc *TestController) RegisterTestRouterGroup(rg *gin.RouterGroup) {
	testroute := rg.Group("/test")
	testroute.POST("/", tc.CreateTest)
	testroute.GET("/getAllTest", tc.GetAllTest)
	testroute.GET("/getTestId", tc.GetTestID)

}
