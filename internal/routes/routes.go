package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhail34/notes-api/internal/handlers"
	"github.com/suhail34/notes-api/internal/middlewares"
)

func SetupRoutes(app *gin.Engine) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/signup", handlers.Connect().SignUpHandler)
		authGroup.POST("/login", handlers.Connect().LoginHandler)
	}
	noteGroup := app.Group("/api/notes")
	noteGroup.Use(middlewares.AuthMiddleware())
	{
		noteGroup.POST("/", handlers.Connect().CreateNotes)
    noteGroup.GET("/", handlers.Connect().GetAllNotesHandler)
    noteGroup.GET("/:id", handlers.Connect().GetNoteHandler)
    noteGroup.PUT("/:id", handlers.Connect().UpdateNoteHandler)
    noteGroup.DELETE("/:id", handlers.Connect().DeleteNoteHandler)
    noteGroup.POST("/:id/share", handlers.Connect().ShareNoteHandler)
    noteGroup.GET("/search", handlers.Connect().SearchTextOnNotes)
	}
}
