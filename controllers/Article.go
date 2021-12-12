package controllers

import (
	"log"
	"net/http"

	models "github.com/atk81/go-blog-post/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetArticle(c *gin.Context) {
	article_id,_ := uuid.Parse(c.Param("article_id"))
	Article, err := models.GetArticle(article_id)
	if err != nil {
		log.Fatalf("An error ocurred %v", err)
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the article.html template
		"article.html",
		// Pass the data that the page uses
		gin.H{
			"title":   Article.Title,
			"payload": Article,
		},
	)
}

func ShowArticleCreationPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the article.html template
		"create-article.html",
		// Pass the data that the page uses
		gin.H{
			"title": "Create article",
		},
	)
}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if id, err := models.CreateArticle(title, content); err != nil {
		log.Fatalf("An error ocurred %v", err)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/article/view/"+id.String())
	}
}
