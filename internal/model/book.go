package model

// Book represents a book in the system.
type Book struct {
	// ID is the unique identifier of the book.
	ID int `json:"id"`

	// Title is the title of the book.
	Title string `json:"title"`

	// Author is the author of the book.
	Author string `json:"author"`

	// PublishedYear is the year the book was published.
	PublishedYear int `json:"publishedYear"`

	// Genre is the genre of the book.
	Genre string `json:"genre"`

	// Summary provides a brief summary or description of the book.
	Summary string `json:"summary"`
}
