package models

// Book represents a book in the database
type Book struct {
	ID              int    `db:"id"               json:"id"`
	Title           string `db:"title"            json:"title"`
	Author          string `db:"author"           json:"author"`
	PublicationYear int    `db:"publication_year" json:"publication_year"`
}
