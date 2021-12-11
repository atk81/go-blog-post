package routes

import (
	"github.com/atk81/go-blog-post/controllers"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.Engine) {
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", controllers.ShowRegistrationPage)

		userRoutes.POST("/register", controllers.Register)

		userRoutes.GET("/login", controllers.ShowLoginPage)
		userRoutes.POST("/login", controllers.PerformLogin)
		userRoutes.GET("/logout", controllers.Logout)
	}
}
