package routes

import (
	"github.com/dondonudon/authentication-authorization/controllers"
	"github.com/dondonudon/authentication-authorization/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
	router.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail)
	router.POST("/forgotpassword", rc.authController.ForgotPassword)
	router.PATCH("/resetpassword/:resetToken", rc.authController.ResetPassword)
	router.GET("/refresh", rc.authController.RefreshAccessToken)

}
