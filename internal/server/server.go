package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/suhail34/notes-api/internal/handlers"
	"github.com/suhail34/notes-api/internal/middlewares"
	"github.com/suhail34/notes-api/internal/routes"
)

func NewServer() *gin.Engine {
	app := gin.Default()
  db, err := handlers.Connect()
  if err != nil {
    log.Fatal(err)
  }
	app.Use(middlewares.RateLimiterMiddleware)
	routes.SetupRoutes(app, db)
	return app
}
