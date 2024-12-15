package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yikuanzz/social/internal/controller"
)

type SocialApiRouter struct {
	userController *controller.UserController
}

func NewSocialApiRouter(uc *controller.UserController) *SocialApiRouter {
	return &SocialApiRouter{
		userController: uc,
	}
}

// Register UnAuthorized API Router
func (s *SocialApiRouter) RegisterUnAuthorizedRoutes(r *gin.RouterGroup) {
	// user
	r.POST("/user/login/email", s.userController.UserLoginByEmail)
	r.POST("/user/register/email", s.userController.UserRegisterByEmail)

}

// Register Authorized API Router
func (router *SocialApiRouter) RegisterAuthorizedApiRouter(r *gin.RouterGroup) {

}
