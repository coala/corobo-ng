package controllers

import (
	"log"
	"net/http"

	"github.com/coala/corobo-ng/services"
	"github.com/coala/corobo-ng/utils"
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
	requestBody, err := utils.GetRequestBody(c)
	if err != nil {
		log.Printf("Aborting, unable to read request body, %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "success": false})
		return
	}

	githubCode := requestBody["code"].(string)
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
		c.SetCookie("token", user.Token, -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"token": user.Token, "success": true})
		return
	}

	c.SetCookie("token", user.Token, 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": user.Token, "success": true})
}

func (base *Controller) GitlabSignUp(c *gin.Context) {
	requestBody, err := utils.GetRequestBody(c)
	if err != nil {
		log.Printf("Aborting, unable to read request body, %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "success": false})
		return
	}

	gitlabCode := requestBody["code"].(string)
	log.Printf("Authenticating user with code %s", gitlabCode)
	userData, err := services.GetGitlabUser(gitlabCode)
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
		c.SetCookie("token", user.Token, -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"token": user.Token, "success": true})
		return
	}

	c.SetCookie("token", user.Token, 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": user.Token, "success": true})
}
