package routes

import (
	"API_Books/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(router *gin.Engine) {
	router.POST("/books", handlers.CreateBook)
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/search", handlers.SearchBooks)
	router.GET("/books/:id", handlers.GetBookByID)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)
}
