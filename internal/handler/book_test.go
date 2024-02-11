package handler

import (
	"book-management/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockStore simulates the behavior of the actual store for testing purposes.
type MockStore struct {
	Books []model.Book
}

// NewMockStore initializes and returns a new instance of MockStore with preloaded data.
func NewMockStore() *MockStore {
	return &MockStore{
		Books: []model.Book{
			{ID: 1, Title: "Test Book 1", Author: "Author 1", PublishedYear: 2021, Genre: "Genre 1", Summary: "Summary 1"},
			{ID: 2, Title: "Test Book 2", Author: "Author 2", PublishedYear: 2022, Genre: "Genre 2", Summary: "Summary 2"},
		},
	}
}

// Implementations of the store interface methods for the MockStore

func (m *MockStore) GetBooks() ([]model.Book, error) {
	return m.Books, nil
}

func (m *MockStore) CreateBook(book *model.Book) error {
	book.ID = len(m.Books) + 1 // Simulate auto-increment
	m.Books = append(m.Books, *book)
	return nil
}

func (m *MockStore) GetBook(id int) (model.Book, error) {
	for _, book := range m.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return model.Book{}, errors.New("book not found")
}

func (m *MockStore) UpdateBook(book *model.Book) error {
	for i, b := range m.Books {
		if b.ID == book.ID {
			m.Books[i] = *book
			return nil
		}
	}
	return errors.New("book not found")
}

func (m *MockStore) DeleteBook(id int) error {
	for i, book := range m.Books {
		if book.ID == id {
			m.Books = append(m.Books[:i], m.Books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

// TestGetBooks tests the GetBooks handler function.
func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewBookHandler(NewMockStore()).GetBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var books []model.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatal(err)
	}

	if len(books) != len(NewMockStore().Books) {
		t.Errorf("Expected %d books, got %d", len(NewMockStore().Books), len(books))
	}
}

// TestCreateBook tests the CreateBook handler function.
func TestCreateBook(t *testing.T) {
	mockStore := NewMockStore()          // Use the mocked store with predefined books
	handler := NewBookHandler(mockStore) // Initialize the handler with the mock store

	newBook := model.Book{
		Title: "New Book", Author: "Author Name", PublishedYear: 2020, Genre: "Fiction", Summary: "A new book summary",
	}
	requestBody, _ := json.Marshal(newBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	http.HandlerFunc(handler.CreateBook).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var bookResponse model.Book
	_ = json.Unmarshal(rr.Body.Bytes(), &bookResponse)

	if bookResponse.Title != newBook.Title {
		t.Errorf("handler returned unexpected body: got title %v want title %v", bookResponse.Title, newBook.Title)
	}
}
