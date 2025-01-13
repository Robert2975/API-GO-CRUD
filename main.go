package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Book represents a book in the database
type Book struct {
	ID              int    `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	Author          string `db:"author" json:"author"`
	PublicationYear int    `db:"publication_year" json:"publication_year"`
}

var db *sqlx.DB

func main() {
	var err error
	// Connect to PostgreSQL
	db, err = sqlx.Connect("postgres", "user=postgres password=123456 dbname=gin_crud sslmode=disable")
	if err != nil {
		log.Fatalln("Falha ao conectar com banco de dados:", err)
	}

	router := gin.Default()

	// Defina os proxies confiáveis (por exemplo, IPs da sua rede)
	err = router.SetTrustedProxies([]string{"192.168.1.1", "192.168.0.46"}) // Substitua pelos IPs confiáveis
	if err != nil {
		panic(err)
	}

	// Routes
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	// Start server
	router.Run(":8080")
}

// Handlers

// Get all books
func getBooks(c *gin.Context) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// Get a single book by ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")
	var book Book
	err := db.Get(&book, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Create a new book
func createBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "INSERT INTO books (title, author, publication_year) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(query, book.Title, book.Author, book.PublicationYear).Scan(&book.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// Update an existing book
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "UPDATE books SET title=$1, author=$2, publication_year=$3 WHERE id=$4"
	_, err := db.Exec(query, book.Title, book.Author, book.PublicationYear, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso!"})
}

// Delete a book
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso!"})
}
