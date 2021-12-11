package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/atk81/go-blog-post/controllers"
)

func UserRegister(router *gin.Engine) {
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", controllers.ShowRegistrationPage)

		userRoutes.POST("/register", controllers.Register)
	}
}
