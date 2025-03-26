package handlers

import (
	"net/http"
	"API_Books/internal/database"
	"API_Books/internal/models"
	"github.com/gin-gonic/gin"
)

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "UPDATE books SET title=$1, author=$2, publication_year=$3 WHERE id=$4"
	_, err := database.DB.Exec(query, book.Title, book.Author, book.PublicationYear, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Livro atualizado com sucesso!"})
}