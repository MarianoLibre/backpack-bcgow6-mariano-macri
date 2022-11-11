package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library.com/internal/domain"
	"library.com/internal/loan"
)

type Loan struct {
	LoanService loan.Service
}

func NewLoan(ls loan.Service) *Loan {
	return &Loan{ls}
}

func (b *Loan) GetAllLoans() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := b.LoanService.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func (b *Loan) GetLoan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		loan, err := b.LoanService.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, loan)
	}
}

func (b *Loan) StoreLoan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.Loan
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		switch {
		case request.BookId == 0:
			ctx.JSON(http.StatusBadRequest, "BookId must be provided")
			return
		case request.UserId == 0:
			ctx.JSON(http.StatusBadRequest, "UserId must be greater than 0")
			return
		}

		id, err := b.LoanService.Save(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		request.Id = id
		ctx.JSON(http.StatusOK, request)
	}
}

func (b *Loan) UpdateLoan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var request domain.Loan
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		switch {
		case request.BookId == 0:
			ctx.JSON(http.StatusBadRequest, "BookId must be provided")
			return
		case request.UserId == 0:
			ctx.JSON(http.StatusBadRequest, "UserId must be greater than 0")
			return
		}

		request.Id = int(id)
		err = b.LoanService.Update(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, request)
	}
}

func (b *Loan) DeleteLoan() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = b.LoanService.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusNoContent, "")
	}
}
