package repository

import (
	"strconv"

	"github.com/prongbang/arch/model"
)

type IBookRepository interface {
	GetAll(page *model.Paging) []*model.Book
	GetById(id int) *model.Book
	Create(b *model.Book) *model.Book
	Update(b *model.Book) *model.Book
	Delete(b *model.Book) *model.Book
}

type BookRepository struct{}

func CreateBookRepository() *BookRepository {
	return &BookRepository{}
}

// implement IBookRepository
func (rep *BookRepository) GetAll(page *model.Paging) []*model.Book {
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

func (rep *BookRepository) GetById(id int) *model.Book {

	return &model.Book{
		ID:    id,
		Name:  "Book " + strconv.Itoa(id),
		Price: float64(1000 * id),
	}
}

func (rep *BookRepository) Create(b *model.Book) *model.Book {

	return b
}

func (rep *BookRepository) Update(b *model.Book) *model.Book {

	return b
}

func (rep *BookRepository) Delete(b *model.Book) *model.Book {

	return b
}
