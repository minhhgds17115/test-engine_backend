package controller

import (
	"net/http"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserServiceImpl
}

func NewController(userServices *services.UserServiceImpl) *UserController {
	return &UserController{
		userService: userServices,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var Candidate models.Candidate
	if err := ctx.ShouldBindJSON(&Candidate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.CreateUser(&Candidate)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUserEmail(ctx *gin.Context) {
	var Contact string = ctx.Param("contact")
	Candidate, err := uc.userService.GetUserEmail(&Contact)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Candidate)
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	Candidate, err := uc.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Candidate)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var Candidate models.Candidate
	if err := ctx.ShouldBindJSON(&Candidate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.UpdateUser(&Candidate)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var FirstName string = ctx.Param("id")
	err := uc.userService.DeleteUser(&FirstName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterRouterGroup(rg *gin.RouterGroup) {
	userroute := rg.Group("/Candidate")
	userroute.GET("/:contact", uc.GetUserEmail)

	userroute.POST("/", uc.CreateUser)
	userroute.PATCH("/:first_name", uc.UpdateUser)
	userroute.DELETE("/:id", uc.DeleteUser)
	userroute.GET("/", uc.GetAllUsers)

}
