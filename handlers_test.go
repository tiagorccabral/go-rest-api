package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBooks(t *testing.T) {

	books = nil

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "LotR", Author: &Author{Firstname: "J.R.R", Lastname: "Tolkien"}})
	books = append(books, Book{ID: "2", Isbn: "845632", Title: "HP", Author: &Author{Firstname: "J.K", Lastname: "Rowling"}})

	t.Run("returns all books", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/api/books", nil)
		response := httptest.NewRecorder()

		// Handler attribution
		handler := http.HandlerFunc(getBooks)

		// Send request
		handler.ServeHTTP(response, request)

		// Test if response got 200 OK
		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Test if body of response contains expected Data
		got := strings.TrimSuffix(response.Body.String(), "\n")
		expected := `[{"id":"1","isbn":"448743","title":"LotR","author":{"firstname":"J.R.R","lastname":"Tolkien"}},{"id":"2","isbn":"845632","title":"HP","author":{"firstname":"J.K","lastname":"Rowling"}}]`
		if got != expected {
			t.Errorf("handler returned wrong response body: got %v want %v",
				got, expected)
		}

	})
}
