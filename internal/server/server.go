package server

import (
	"github.com/gin-gonic/gin"
	"github.com/suhail34/notes-api/internal/middlewares"
	"github.com/suhail34/notes-api/internal/routes"
)

func NewServer() *gin.Engine {
	app := gin.Default()
	app.Use(middlewares.RateLimiterMiddleware)
	routes.SetupRoutes(app)
	return app
}
