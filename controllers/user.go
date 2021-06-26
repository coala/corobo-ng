package controllers

import (
	"log"
	"net/http"

	"github.com/coala/corobo-ng/services"
	"github.com/gin-gonic/gin"
)

func (base *Controller) GetUser(c *gin.Context) {
	user, err := services.GetUser(base.DB, c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error while retrieving user", "success": false})
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}

func (base *Controller) GithubSignUp(c *gin.Context) {
	githubCode := c.Request.URL.Query().Get("code")
	log.Printf("Authenticating user with code %s", githubCode)

	userData, err := services.GetGithubUser(githubCode)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error while retrieving user", "success": false})
		return
	}

	log.Printf("Got user data, %v", userData)

	user, err := services.GetUserByEmail(base.DB, userData["email"].(string))
	if err != nil {
		// Create a new user as it does not exist
		user, err = services.CreateUser(base.DB, userData)
		if err != nil {
			log.Printf("Failed to create new user, %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error while creating user", "success": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": user.Token, "success": true})
	}

	c.JSON(http.StatusOK, gin.H{"token": user.Token, "success": true})
}
