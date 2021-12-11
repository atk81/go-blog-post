package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/atk81/go-blog-post/controllers"
)

func ArticlesRegister(router *gin.Engine) {
	router.GET("/",controllers.ShowIndexPage)
}