package server

import (
	"github.com/coala/corobo-ng/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	// Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())

	api := controllers.Controller{DB: db}

	// Endpoint to check health through /health
	router.GET("/health", api.Health)

	// All endpoints realted to /user
	userGroup := router.Group("user")
	{
		userGroup.GET("/:id", api.GetUser)
	}

	// Endpoints related to the auth flow
	signupGroup := router.Group("login")
	{
		signupGroup.POST("/", api.GithubSignUp)
	}

	return router
}
