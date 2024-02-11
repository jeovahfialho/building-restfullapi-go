package store

import (
	"database/sql"

	"book-management/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

// Store defines the interface for interacting with the book store.
// Any type that implements these methods can be used as a store for book management.
type Store interface {
	// CreateBook inserts a new book record into the store.
	CreateBook(book *model.Book) error

	// GetBooks retrieves all books from the store.
	GetBooks() ([]model.Book, error)

	// GetBook retrieves a single book by its ID from the store.
	GetBook(id int) (model.Book, error)

	// UpdateBook updates an existing book record in the store.
	UpdateBook(book *model.Book) error

	// DeleteBook deletes a book record from the store by its ID.
	DeleteBook(id int) error
}

// DBStore represents a database implementation of the Store interface.
type DBStore struct {
	db *sql.DB
}

// NewDB creates a new instance of DBStore with the provided data source name (DSN).
// It establishes a connection to the database and returns a DBStore object.
func NewDB(dsn string) (*DBStore, error) {
	// Open a new database connection.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Ping the database to check if the connection is valid.
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// Return a new instance of DBStore.
	return &DBStore{db: db}, nil
}

// Close closes the database connection.
func (store *DBStore) Close() error {
	return store.db.Close()
}

// CreateBook inserts a new book record into the database.
func (store *DBStore) CreateBook(book *model.Book) error {
	_, err := store.db.Exec(`INSERT INTO books (title, author, publishedYear, genre, summary) VALUES (?, ?, ?, ?, ?)`,
		book.Title, book.Author, book.PublishedYear, book.Genre, book.Summary)
	return err
}

// GetBooks retrieves all books from the database.
func (store *DBStore) GetBooks() ([]model.Book, error) {
	rows, err := store.db.Query(`SELECT id, title, author, publishedYear, genre, summary FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Genre, &book.Summary); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// GetBook retrieves a single book by its ID from the database.
func (store *DBStore) GetBook(id int) (model.Book, error) {
	var book model.Book
	row := store.db.QueryRow(`SELECT id, title, author, publishedYear, genre, summary FROM books WHERE id = ?`, id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Genre, &book.Summary); err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// UpdateBook updates an existing book record in the database.
func (store *DBStore) UpdateBook(book *model.Book) error {
	_, err := store.db.Exec(`UPDATE books SET title = ?, author = ?, publishedYear = ?, genre = ?, summary = ? WHERE id = ?`,
		book.Title, book.Author, book.PublishedYear, book.Genre, book.Summary, book.ID)
	return err
}

// DeleteBook deletes a book record from the database by its ID.
func (store *DBStore) DeleteBook(id int) error {
	_, err := store.db.Exec(`DELETE FROM books WHERE id = ?`, id)
	return err
}
