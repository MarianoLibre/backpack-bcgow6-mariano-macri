package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

type Request struct {
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(*c)
		if err != nil {
			web.Error(c, http.StatusBadGateway, "%s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		seller, err := s.sellerService.GetSellerById(*c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, seller)
	}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := c.Bind(&request); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := Validate(request); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		var seller domain.Seller
		if request.Address != "" && request.CompanyName != "" && request.Telephone != "" && request.CID != 0 {
			seller.CID = request.CID
			seller.Address = request.Address
			seller.CompanyName = request.CompanyName
			seller.Telephone = request.Telephone
		} else {
			web.Error(c, http.StatusBadRequest, "None of the attributes can be empty or zero")
			return
		}
		seller, err := s.sellerService.Save(*c, seller)
		if err != nil {
			web.Error(c, http.StatusConflict, "%s", err.Error())
			return
		}
		web.Success(c, http.StatusCreated, seller)
	}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := c.Bind(&request); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := Validate(request); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		var newSeller domain.Seller
		newSeller.CID = request.CID
		newSeller.Address = request.Address
		newSeller.CompanyName = request.CompanyName
		newSeller.Telephone = request.Telephone
		seller, err := s.sellerService.Update(*c, newSeller)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		web.Success(c, http.StatusOK, seller)
	}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.sellerService.Delete(*c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err.Error())
			return
		}
		web.Success(c, http.StatusNoContent, "")
	}
}

func Validate(s Request) error {
	validate := validator.New()
	return validate.Struct(s)
}
