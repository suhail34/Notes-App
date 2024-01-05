package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhail34/notes-api/internal/handlers"
	"github.com/suhail34/notes-api/internal/middlewares"
)

func SetupRoutes(app *gin.Engine, db *handlers.DB) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/signup", db.SignUpHandler)
		authGroup.POST("/login", db.LoginHandler)
	}
	noteGroup := app.Group("/api/notes")
	noteGroup.Use(middlewares.AuthMiddleware())
	{
		noteGroup.POST("/", db.CreateNotes)
    noteGroup.GET("/", db.GetAllNotesHandler)
    noteGroup.GET("/:id", db.GetNoteHandler)
    noteGroup.PUT("/:id", db.UpdateNoteHandler)
    noteGroup.DELETE("/:id", db.DeleteNoteHandler)
    noteGroup.POST("/:id/share", db.ShareNoteHandler)
    noteGroup.GET("/search", db.SearchTextOnNotes)
	}
}
