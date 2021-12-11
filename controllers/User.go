package controllers

import (
	"net/http"

	model "github.com/atk81/go-blog-post/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func generateSessionToken() string {
	// Create a new random session token
	sessionToken := uuid.New().String()
	return sessionToken
}

func ShowRegistrationPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the register.html template
		"register.html",
		// Pass the data that the page uses
		gin.H{
			"title": "Register",
		},
	)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if err := model.RegisterNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		c.HTML(
			http.StatusOK,
			"login-successful.html",
			gin.H{
				"title": "Login successful",
			},
		)
	} else{
		// If the username/password combination is invalid,
        // show the error message on the login page
        c.HTML(http.StatusBadRequest, "register.html", gin.H{
            "ErrorTitle":   "Registration Failed",
            "ErrorMessage": err.Error()})

	}
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the login.html template
		"login.html",
		// Pass the data that the page uses
		gin.H{
			"title": "Log In",
		},
	)
}

func PerformLogin(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    if model.IsUserValid(username, password) {
        token := generateSessionToken()
        c.SetCookie("token", token, 3600, "", "", false, true)

		c.HTML(
			http.StatusOK,
			"login-successful.html",
			gin.H{
				"title": "Login successful",
			},
		)

    } else {
        c.HTML(http.StatusBadRequest, "login.html", gin.H{
            "ErrorTitle":   "Login Failed",
            "ErrorMessage": "Invalid credentials provided"})
    }
}

func Logout(c *gin.Context) {
    c.SetCookie("token", "", -1, "", "", false, true)

    c.Redirect(http.StatusTemporaryRedirect, "/")
}
