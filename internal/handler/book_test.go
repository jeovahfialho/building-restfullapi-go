package handler

import (
	"book-management/internal/model"
	"book-management/internal/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

type MockStore struct {
	store.Store
}

func (m *MockStore) GetBooks() ([]model.Book, error) {
	books := []model.Book{
		{ID: 1, Title: "Test Book 1", Author: "Author 1", PublishedYear: 2021, Genre: "Genre 1", Summary: "Summary 1"},
		{ID: 2, Title: "Test Book 2", Author: "Author 2", PublishedYear: 2022, Genre: "Genre 2", Summary: "Summary 2"},
	}
	return books, nil
}

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewBookHandler(&MockStore{}).GetBooks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var books []model.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatal(err)
	}

	if len(books) != 2 {
		t.Errorf("Expected 2 books, got %d", len(books))
	}
}

// TestCreateBook tests the CreateBook handler function.
func TestCreateBook(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TestCreateBook panicked: %v", r)
		}
	}()
	// Define a sample book to be created.
	newBook := model.Book{
		ID:            3,
		Title:         "New Test Book",
		Author:        "New Author",
		PublishedYear: 2023,
		Genre:         "New Genre",
		Summary:       "New Summary",
	}

	// Marshal the book into JSON format.
	reqBody, err := json.Marshal(newBook)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the book data in the body.
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a new BookHandler instance with the MockStore and invoke CreateBook handler function.
	handler := http.HandlerFunc(NewBookHandler(&MockStore{}).CreateBook)
	handler.ServeHTTP(rr, req)

	// Check if the status code returned by the handler is as expected.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Decode the response body to check if the created book matches the expected one.
	var createdBook model.Book
	err = json.Unmarshal(rr.Body.Bytes(), &createdBook)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the created book matches the expected one.
	if !reflect.DeepEqual(createdBook, newBook) {
		t.Errorf("Expected created book to match input book, got %+v", createdBook)
	}
}

// TestGetBook tests the GetBook handler function.
func TestGetBook(t *testing.T) {
	// Define the ID of the book to retrieve.
	bookID := 2

	// Create a new HTTP request targeting the GetBook handler with the book ID.
	req, err := http.NewRequest("GET", "/books/"+strconv.Itoa(bookID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a new BookHandler instance with the MockStore and invoke GetBook handler function.
	handler := http.HandlerFunc(NewBookHandler(&MockStore{}).GetBook)
	handler.ServeHTTP(rr, req)

	// Check if the status code returned by the handler is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body into a book object.
	var retrievedBook model.Book
	err = json.Unmarshal(rr.Body.Bytes(), &retrievedBook)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the retrieved book matches the expected one with the specified ID.
	if retrievedBook.ID != bookID {
		t.Errorf("Expected book with ID %d, got %+v", bookID, retrievedBook)
	}
}

// TestUpdateBook tests the UpdateBook handler function.
func TestUpdateBook(t *testing.T) {
	// Define the ID of the book to update.
	bookID := 1

	// Define updated book data.
	updatedBook := model.Book{
		ID:            bookID,
		Title:         "Updated Test Book",
		Author:        "Updated Author",
		PublishedYear: 2024,
		Genre:         "Updated Genre",
		Summary:       "Updated Summary",
	}

	// Marshal the updated book into JSON format.
	reqBody, err := json.Marshal(updatedBook)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the updated book data in the body.
	req, err := http.NewRequest("PUT", "/books/"+strconv.Itoa(bookID), bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a new BookHandler instance with the MockStore and invoke UpdateBook handler function.
	handler := http.HandlerFunc(NewBookHandler(&MockStore{}).UpdateBook)
	handler.ServeHTTP(rr, req)

	// Check if the status code returned by the handler is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body to check if the updated book matches the expected one.
	var updatedResponse model.Book
	err = json.Unmarshal(rr.Body.Bytes(), &updatedResponse)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the updated book matches the expected one.
	if !reflect.DeepEqual(updatedResponse, updatedBook) {
		t.Errorf("Expected updated book to match input book, got %+v", updatedResponse)
	}
}

// TestDeleteBook tests the DeleteBook handler function.
func TestDeleteBook(t *testing.T) {
	// Define the ID of the book to delete.
	bookID := 1

	// Create a new HTTP request targeting the DeleteBook handler with the book ID.
	req, err := http.NewRequest("DELETE", "/books/"+strconv.Itoa(bookID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response.
	rr := httptest.NewRecorder()

	// Create a new BookHandler instance with the MockStore and invoke DeleteBook handler function.
	handler := http.HandlerFunc(NewBookHandler(&MockStore{}).DeleteBook)
	handler.ServeHTTP(rr, req)

	// Check if the status code returned by the handler is as expected.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body to check if the response message matches the expected one.
	var res map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the response message matches the expected one.
	expectedMessage := "Book deleted successfully"
	if res["message"] != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, res["message"])
	}
}
