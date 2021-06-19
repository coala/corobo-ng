package controllers

import (
	"net/http"

	"github.com/coala/corobo-ng/models"
	"github.com/coala/corobo-ng/services"
	"github.com/gin-gonic/gin"
)

var userModel = new(models.User)

func (base *Controller) GetUser(c *gin.Context) {
	user, err := services.GetUser(base.DB, c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error while retrieving user", "success": false})
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}

func (base *Controller) GithubSignUp(c *gin.Context) {
	githubToken := c.Params.ByName("code")

	user, err := services.GetGithubUser(githubToken)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error while retrieving user", "success": false})
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "success": true})
}
