package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *Controller) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "healthy",
	})
}
