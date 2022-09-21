package controller

import (
	"fmt"
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

// CandidateInformation
func (tc *TestController) StoreCandidateInfo(ctx *gin.Context) {
	var candidateInformation models.CandidateInformation
	
	if err := ctx.ShouldBindJSON(&candidateInformation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println("Global ")
		return
	}

	fmt.Println("global handler ", candidateInformation)

	err := tc.testService.StoreCandidateInfo(&candidateInformation)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		fmt.Println("Information return ")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})

}

// Return Answer

func (tc *TestController) ReturnAnswer(ctx *gin.Context) {
	var returnedAnswer models.ReturnedAnswer

	if err := ctx.ShouldBindJSON(&returnedAnswer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := tc.testService.ReturnAnswer(&returnedAnswer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// func (tc *TestController) GetInfo(ctx *gin.Context) {
// 	testId, _ := strconv.Atoi(ctx.Param("id"))
// 	Candidate, err := tc.testService.GetTestID(&testId)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, Candidate)
// }

func (tc *TestController) GetTestID(ctx *gin.Context) {
	testId, _ := strconv.Atoi(ctx.Param("id"))
	Test, err := tc.testService.GetTestID(&testId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Test)
}

func (tc *TestController) RegisterTestRouterGroup(rg *gin.RouterGroup) {
	testroute := rg.Group("/test")
	testroute.PATCH("/:id", tc.UpdateTest)
	testroute.GET("", tc.GetAllTest)
	testroute.GET("/:id", tc.GetTestID)
	testroute.POST("/candidate-info", tc.StoreCandidateInfo)
	testroute.POST("/answer", tc.ReturnAnswer)

}
