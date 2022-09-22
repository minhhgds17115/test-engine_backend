package controller

import (
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

type CandidateController struct {
	userService *services.CandidateServiceImpl
}

func NewController(userServices *services.CandidateServiceImpl) *CandidateController {
	return &CandidateController{
		userService: userServices,
	}
}

func (uc *CandidateController) CreateCandidate(ctx *gin.Context) {
	var Candidate models.Candidate
	if err := ctx.ShouldBindJSON(&Candidate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.CreateCandidate(&Candidate)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CandidateController) GetCandidateEmail(ctx *gin.Context) {
	var Contact string = ctx.Param("contact")
	Candidate, err := uc.userService.GetCandidateEmail(&Contact)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Candidate)
}

func (uc *CandidateController) GetAllCandidates(ctx *gin.Context) {
	Candidate, err := uc.userService.GetAllCandidates()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Candidate)
}

func (uc *CandidateController) UpdateCandidate(ctx *gin.Context) {
	var Candidate models.Candidate
	if err := ctx.ShouldBindJSON(&Candidate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.UpdateCandidate(&Candidate)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CandidateController) DeleteCandidate(ctx *gin.Context) {
	var FirstName string = ctx.Param("id")
	err := uc.userService.DeleteCandidate(&FirstName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CandidateController) CandidateInformation(ctx *gin.Context) {
	var userInformation models.CandidateInformation
	if err := ctx.ShouldBindJSON(&userInformation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.CandidateInformation(&userInformation)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CandidateController) GetCandidateTestID(ctx *gin.Context) {
	testId, _ := strconv.Atoi(ctx.Param("id"))
	CandidateInformation, err := uc.userService.GetCandidateTestID(&testId)
	if CandidateInformation == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, CandidateInformation)

}

func (uc *CandidateController) StoreTestCandidate(ctx *gin.Context) {
	var testInfo models.Test

	if err := ctx.ShouldBindJSON(&testInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.userService.StoreTestCandidate(&testInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *CandidateController) RegisterRouterGroup(rg *gin.RouterGroup) {
	userroute := rg.Group("/Candidate")
	// userroute.GET("/:contact", uc.GetCandidateEmail)
	userroute.GET("/:id", uc.GetCandidateTestID)
	userroute.POST("/user-information", uc.CandidateInformation)
	userroute.POST("/", uc.CreateCandidate)
	userroute.PATCH("/:firstname", uc.UpdateCandidate)
	userroute.DELETE("/:id", uc.DeleteCandidate)
	userroute.GET("/", uc.GetAllCandidates)

	userroute.POST("/store-candidate", uc.StoreTestCandidate)

}
