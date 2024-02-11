package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"book-management/internal/model"
	"book-management/internal/store"

	"github.com/gorilla/mux"
)

// BookHandler handles HTTP requests related to book management.
type BookHandler struct {
	// Store is the interface for interacting with the book store.
	Store store.Store
}

// NewBookHandler creates a new instance of BookHandler with the provided store.
func NewBookHandler(store store.Store) *BookHandler {
	return &BookHandler{Store: store}
}

// CreateBook handles the creation of a new book.
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Store.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBooks handles the retrieval of all books.
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Store.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// GetBook handles the retrieval of a single book by its ID.
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := h.Store.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles the update of an existing book.
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = id
	err = h.Store.UpdateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// DeleteBook handles the deletion of a book by its ID.
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.Store.DeleteBook(id)
	if err != nil {
		// Assuming err could be a not found error or an internal server error,
		// you might want to differentiate here in a real application.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Setting the header to application/json for the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // or http.StatusNoContent if you prefer no content

	// Creating a response struct
	type response struct {
		Message string `json:"message"`
	}

	res := response{
		Message: "Book deleted successfully",
	}

	json.NewEncoder(w).Encode(res)
}
