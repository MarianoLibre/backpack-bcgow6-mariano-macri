package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library.com/internal/domain"
	"library.com/internal/library"
)

type Library struct {
	LibraryService library.Service
}

func NewLibrary(ls library.Service) *Library {
	return &Library{ls}
}

func (l *Library) GetAllLibraries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := l.LibraryService.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func (l *Library) GetLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		book, err := l.LibraryService.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, book)
	}
}

func (l *Library) StoreLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.Library
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		id, err := l.LibraryService.Save(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		request.Id = id
		ctx.JSON(http.StatusOK, request)
	}
}

func (l *Library) UpdateLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var request domain.Library
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		request.Id = int(id)
		err = l.LibraryService.Update(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, request)
	}
}

func (l *Library) DeleteLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = l.LibraryService.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	}
}
