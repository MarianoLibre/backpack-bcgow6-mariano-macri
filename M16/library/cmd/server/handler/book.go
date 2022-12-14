package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library.com/internal/book"
	"library.com/internal/domain"
)

type Book struct {
	BookService book.Service
}

func NewBook(bs book.Service) *Book {
	return &Book{bs}
}

func (b *Book) GetAllBooks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := b.BookService.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func (b *Book) GetBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		book, err := b.BookService.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, book)
	}
}

func (b *Book) StoreBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.Book
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		id, err := b.BookService.Save(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		request.Id = id
		ctx.JSON(http.StatusOK, request)
	}
}

func (b *Book) UpdateBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var request domain.Book
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		request.Id = int(id)
		err = b.BookService.Update(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, request)
	}
}

func (b *Book) DeleteBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = b.BookService.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	}
}
