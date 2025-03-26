package handlers

import (
	"net/http"
	"API_Books/internal/database"
	"API_Books/internal/models"

	"github.com/gin-gonic/gin"
)


// GetBookByID handles GET /books/:id
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	err := database.DB.Get(&book, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}