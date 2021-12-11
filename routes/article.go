package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/atk81/go-blog-post/controllers"
)

func ArticleRegister(router *gin.Engine){
	router.GET("/article/view/:article_id",controllers.GetArticle)
}