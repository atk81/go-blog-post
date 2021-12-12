package controllers

import (
	"log"
	"net/http"

	models "github.com/atk81/go-blog-post/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	Articles,err := models.GetAllArticles()
	if err!= nil{
		log.Fatalf("An error ocurred %v",err)
	}
	
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": Articles,
			"is_logged_in": c.GetBool("is_logged_in"),
		},
	)
}
