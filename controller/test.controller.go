package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"example.com/m/v2/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type TestController struct {
	testService *services.TestServiceImpl
}

var Validate *validator.Validate

func NewTestController(testService *services.TestServiceImpl) *TestController {
	return &TestController{
		testService: testService,
	}
}

func init() {
	Validate = validator.New()
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

	if err := Validate.Struct(candidateInformation); err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			return
		}
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

	if err := Validate.Struct(&returnedAnswer); err != nil {
		ctx.JSON(http.StatusExpectationFailed, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get Test by ID
func (tc *TestController) GetTestID(ctx *gin.Context) {

	testId, _ := strconv.Atoi(ctx.Param("id"))
	ReturnedAnswer, err := tc.testService.GetTestID(&testId)
	fmt.Println("Begin validation")
	if ReturnedAnswer == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	if err := Validate.Struct(ReturnedAnswer); err != nil {
		if _, ok := err.(*validator.ValidationErrors); ok {
			return
		}
	}

	ctx.JSON(http.StatusOK, ReturnedAnswer)

}

func (tc *TestController) RegisterTestRouterGroup(rg *gin.RouterGroup) {
	testroute := rg.Group("/test")
	testroute.PATCH("/:id", tc.UpdateTest)
	testroute.GET("", tc.GetAllTest)
	testroute.GET("/:id", tc.GetTestID)
	testroute.POST("/candidate-info", tc.StoreCandidateInfo)
	testroute.POST("/answer", tc.ReturnAnswer)

}
