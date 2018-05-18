package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/prongbang/arch/model"
	"github.com/prongbang/arch/repository"
)

type BookHandler struct {
	Repo repository.IBookRepository
}

func (h *BookHandler) GetAll(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	if page != "" && limit != "" {
		pageInt, pErr := strconv.Atoi(page)
		limitInt, lErr := strconv.Atoi(limit)
		if pageInt > 0 && limitInt > 0 && pErr == nil && lErr == nil {
			books := h.Repo.GetAll(&model.Paging{
				Page:  pageInt,
				Limit: limitInt,
			})
			return c.JSON(http.StatusOK, books)
		}
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Please enter parameter page > 0 and limit > 0",
		})
	}
	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": "Required parameter page and limit",
	})
}

func (h *BookHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		idInt, iErr := strconv.Atoi(id)
		if idInt > 0 && iErr == nil {
			book := h.Repo.GetById(idInt)
			return c.JSON(http.StatusOK, book)
		}
	}
	return c.JSON(http.StatusBadGateway, echo.Map{
		"message": "Required parameter id",
	})
}

func (h *BookHandler) Create(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Data not found!",
		})
	}
	return c.JSON(http.StatusCreated, h.Repo.Create(&book))
}

func (h *BookHandler) Update(c echo.Context) error {
	var book model.Book
	id := c.Param("id")
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Data not found!",
			"error":   err,
		})
	}

	if id != "" {
		if idInt, err := strconv.Atoi(id); idInt > 0 && err == nil {
			var result *model.Book
			book.ID = idInt
			switch c.Request().Method {
			case "PUT":
				check := h.Repo.GetById(idInt)
				if check.ID > 0 {
					result = h.Repo.Update(&book)
				} else {
					result = h.Repo.Create(&book)
				}
				break
			case "PATCH":
				result = h.Repo.Update(&book)
				break
			}
			return c.JSON(http.StatusOK, result)
		}
	}

	return c.JSON(http.StatusBadGateway, echo.Map{
		"message": "Required parameter id",
	})
}

func (h *BookHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		if idInt, err := strconv.Atoi(id); idInt > 0 && err == nil {
			return c.JSON(http.StatusOK, h.Repo.Delete(&model.Book{
				ID:    idInt,
				Name:  "Book " + id,
				Price: float64(1000 * idInt),
			}))
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{
		"message": "Id not found!",
	})
}
