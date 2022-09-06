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
	// user.FirstName = strings.TrimSpace(user.FirstName)

	// if user.FirstName == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
	// 	return
	// }
	// user.LastName = strings.TrimSpace(user.LastName)
	// if user.LastName == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
	// 	return
	// }

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

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.userService.GetAllUsers()
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
	var FirstName string = ctx.Param("first_name")
	err := uc.userService.DeleteUser(&FirstName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterRouterGroup(rg *gin.RouterGroup) {
	userroute := rg.Group("/users")
	userroute.GET("/GetUserEmail/:email", uc.GetUserEmail)
	userroute.POST("/CreateUser/:id", uc.CreateUser)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:FirstName", uc.DeleteUser)
	userroute.GET("/GetAllUsers", uc.GetAllUsers)

}
