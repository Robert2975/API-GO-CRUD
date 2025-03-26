package handlers

import (
	"net/http"
	"API_Books/internal/database"
	"github.com/gin-gonic/gin"
)


// DeleteBook handles DELETE /books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso!"})
}