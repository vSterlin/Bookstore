package book

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockBookRepo struct{}

var mockBooks = []*Book{
	{Id: 1, Author: "Aldous Huxley", Title: "Brave New World", Pages: 311, PublishDate: "1932"},
	{Id: 2, Author: "George Orwell", Title: "1984", Pages: 328, PublishDate: "1949"},
}

func (m *mockBookRepo) GetMany() []*Book {
	return mockBooks
}

func (m *mockBookRepo) GetOne(id int) *Book {
	return mockBooks[id-1]
}

func TestGetManyController(t *testing.T) {

	br := &mockBookRepo{}
	bs := NewBookService(br)
	bc := NewBookController(bs)

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	bc.GetBooks(w, req)
	res := w.Result()
	defer res.Body.Close()

	books := []*Book{}

	json.NewDecoder(res.Body).Decode(&books)

	if books[0].Title != mockBooks[0].Title {
		t.Errorf("Expected Title to equal \"%s\", got \"%s\" instead from %+v\n", mockBooks[0].Title, books[0].Title, books[0])
	}

	if len(books) != len(mockBooks) {
		t.Errorf("Expected length of response to be %d, got %d instead\n", len(mockBooks), len(books))
	}
}

func TestGetOneController(t *testing.T) {

}
