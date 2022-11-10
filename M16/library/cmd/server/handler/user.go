package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library.com/internal/domain"
	"library.com/internal/user"
)

type User struct {
	UserSevice user.Service
}

func NewUser(us user.Service) *User {
	return &User{us}
}

func (u *User) GetAllLibraries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := u.UserSevice.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func (u *User) GetLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		book, err := u.UserSevice.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, book)
	}
}

func (u *User) StoreLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.User
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		switch {
		case request.Name == "":
			ctx.JSON(http.StatusBadRequest, "Name must be provided")
			return
		case request.Age == 0:
			ctx.JSON(http.StatusBadRequest, "Age must be provided")
			return
		}

		id, err := u.UserSevice.Save(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		request.Id = id
		ctx.JSON(http.StatusOK, request)
	}
}

func (u *User) UpdateLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var request domain.User
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		switch {
		case request.Name == "":
			ctx.JSON(http.StatusBadRequest, "Name must be provided")
			return
		case request.Age == 0:
			ctx.JSON(http.StatusBadRequest, "Age must be provided")
			return
		}

		request.Id = int(id)
		err = u.UserSevice.Update(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, request)
	}
}

func (u *User) DeleteLibrary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = u.UserSevice.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	}
}
