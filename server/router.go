package server

import (
	"github.com/coala/corobo-ng/controllers"
	"github.com/coala/corobo-ng/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	cors := cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
	})

	// Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors)

	api := controllers.Controller{DB: db}

	// Endpoint to check health through /health
	router.GET("/health", api.Health)

	// Invalidate token and logout
	router.PUT("/logout", api.LogoutUser)

	// All endpoints realted to /user
	userGroup := router.Group("user").Use(middlewares.AuthMiddleware(db))
	{
		userGroup.GET("/:id", api.GetUser)
	}

	// Endpoints related to the auth flow
	signupGroup := router.Group("login")
	{
		signupGroup.POST("/github", api.GithubSignUp)
		signupGroup.POST("/gitlab", api.GitlabSignUp)
	}

	return router
}
