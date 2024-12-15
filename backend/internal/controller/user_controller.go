package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yikuanzz/social/internal/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{
		service: us,
	}
}

// User login by email
func (uc *UserController) UserLoginByEmail(ctx *gin.Context) {

}

// User regiser by email
func (uc *UserController) UserRegisterByEmail(ctx *gin.Context) {

}
