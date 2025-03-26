package handlers

import (
	"API_Books/internal/database"
	"API_Books/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBook handles POST /books
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se o livro(titulo) já existe no banco de dados.
	//A função EXISTS retorna TRUE se a subconsulta dentro dela retornar
	// pelo menos uma linha e FALSE caso contrário.
	var exists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM books WHERE title = $1)"
	err := database.DB.QueryRow(checkQuery, book.Title).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Livro já foi cadastrado"})
		return
	}

	query := "INSERT INTO books (title, author, publication_year) VALUES ($1, $2, $3) RETURNING id"
	err = database.DB.QueryRow(query, book.Title, book.Author, book.PublicationYear).Scan(&book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}
