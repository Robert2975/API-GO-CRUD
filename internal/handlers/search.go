package handlers

import (
	"net/http"
	"API_Books/internal/database"
	"API_Books/internal/models"

	"github.com/gin-gonic/gin"
)

func SearchBooks(c *gin.Context) {
	var books []models.Book
	search := c.Query("search")
	err := database.DB.Select(&books, "SELECT * FROM books WHERE title LIKE $1", "%"+search+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}