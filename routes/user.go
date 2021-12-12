package routes

import (
	"github.com/atk81/go-blog-post/controllers"
	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.Engine) {
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", controllers.EnsureNotLoggedIn(), controllers.ShowRegistrationPage)

		userRoutes.POST("/register", controllers.EnsureNotLoggedIn(), controllers.Register)

		userRoutes.GET("/login", controllers.EnsureNotLoggedIn(), controllers.ShowLoginPage)
		userRoutes.POST("/login", controllers.EnsureNotLoggedIn(), controllers.PerformLogin)
		userRoutes.GET("/logout", controllers.EnsureLoggedIn(), controllers.Logout)
	}
}
