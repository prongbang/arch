package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/prongbang/arch/handler"
	"github.com/prongbang/arch/repository"
	"github.com/stretchr/testify/assert"
)

const (
	bookJson = "{\"id\":1,\"name\":\"Golang\",\"price\":123.321}"
)

func TestCreateBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/book", strings.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.BookHandler{repository.CreateBookRepository()}

	// Assertions
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, bookJson, rec.Body.String())
	}
}

func TestGetAllByPaged(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/books?page=1&limit=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.BookHandler{repository.CreateBookRepository()}

	// Assertions
	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body)
	}
}

func TestPutUpdate(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/book/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := handler.BookHandler{repository.CreateBookRepository()}

	// Assertions
	if assert.NoError(t, h.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bookJson, rec.Body.String())
	}
}

// TODO ...
