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
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUserEmail(ctx *gin.Context) {
	var email string = ctx.Param("email")
	user, err := uc.userService.GetUserEmail(&email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.userService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var firstnames string = ctx.Param("firstnames")
	err := uc.userService.DeleteUser(&firstnames)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterRouterGroup(rg *gin.RouterGroup) {
	userroute := rg.Group("/users")
	userroute.GET("/:id", uc.GetUserEmail)
	userroute.POST("/", uc.CreateUser)
	userroute.PATCH("/", uc.UpdateUser)
	userroute.DELETE("/", uc.DeleteUser)
	userroute.GET("/all", uc.GetAll)

}
