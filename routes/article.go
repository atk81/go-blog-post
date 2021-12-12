package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/atk81/go-blog-post/controllers"
)

func ArticleRegister(router *gin.Engine){
	articleRoutes := router.Group("/article")
    {
        // route from Part 1 of the tutorial
        articleRoutes.GET("/view/:article_id", controllers.GetArticle)

        articleRoutes.GET("/create", controllers.EnsureLoggedIn(), controllers.ShowArticleCreationPage)

        articleRoutes.POST("/create", controllers.EnsureLoggedIn(), controllers.CreateArticle)
    }
}