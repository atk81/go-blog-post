package routes

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {
	// Register all the routes here.
	ArticlesRegister(router)
	ArticleRegister(router)
}
