package main

import (
	routes "github.com/atk81/go-blog-post/routes"
	utils "github.com/atk81/go-blog-post/utils"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	utils.CreateDBConnection()
	router = gin.Default()
	// Initilize Routes
	routes.Router(router)
	router.LoadHTMLGlob("templates/*.html")
	router.Run(":8080")
}
