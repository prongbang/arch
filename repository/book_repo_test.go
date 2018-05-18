package repository_test

import (
	"strconv"
	"testing"

	"github.com/prongbang/arch/model"
)

type MockBookRepository struct{}

func (repo *MockBookRepository) GetAll(page *model.Paging) []*model.Book {
	start := (page.Page * page.Limit) - page.Limit
	size := 1000
	var m []*model.Book
	for i := start + 1; i <= size && i <= (page.Page*page.Limit); i++ {
		m = append(m, &model.Book{
			ID:    i,
			Name:  "Book " + strconv.Itoa(i),
			Price: float64(size * i),
		})
	}
	return m
}

func (repo *MockBookRepository) GetById(id int) *model.Book {
	return &model.Book{
		ID:    id,
		Name:  "Book " + strconv.Itoa(id),
		Price: float64(1000 * id),
	}
}

func (repo *MockBookRepository) Create(b *model.Book) *model.Book {

	return b
}

func (repo *MockBookRepository) Update(b *model.Book) *model.Book {

	return b
}

func (repo *MockBookRepository) Delete(b *model.Book) *model.Book {

	return b
}

func TestGetAllByPeged(t *testing.T) {
	repo := &MockBookRepository{}
	books := repo.GetAll(&model.Paging{1, 10})
	if len(books) != 10 {
		t.Fatalf("Expected 10 but got %d", len(books))
	}
}
