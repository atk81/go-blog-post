package routes

import (
	"github.com/atk81/go-blog-post/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	// Add New Middleware to ensure authentication
	router.Use(controllers.SetUserStatus())
	// Register all the routes here.
	ArticlesRegister(router)
	ArticleRegister(router)
	UserRegister(router)
}
