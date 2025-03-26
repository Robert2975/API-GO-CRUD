package handlers

import (
	"API_Books/internal/database"
	"API_Books/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks handles GET /books
func GetBooks(c *gin.Context) {
	var books []models.Book
	err := database.DB.Select(&books, "SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}
