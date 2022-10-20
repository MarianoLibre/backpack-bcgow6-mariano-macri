package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

// BuyerById 	godoc
// @Summary 	Get buyer by id
// @Tags 		Buyers
// @Description get buyer by id
// @Accept  	json
// @Produce  	json
// @Param 		token header int true "token"
// @Param 		id path int true "id"
// @Success 200 	{object} 	web.response
// @Failure 401 	{object} 	web.response 	"Buyer not exist"
// @Failure 404 	{object} 	web.response 	"Buyer not found"
// @Router /buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error()) //400 - Ver bien el codigo de error
			return
		}

		buyer, err := b.buyerService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error()) //400
			return
		}

		web.Success(c, http.StatusOK, buyer) 
	}
}

// ListBuyers godoc
// @Summary		List buyers
// @Tags 		Buyers
// @Description get buyers
// @Accept  	json
// @Produce  	json
// @Param 		token header int true "token"
// @Success 200 	{object} 	web.response
// @Success 204 	{object}    web.response 	"empty buyer list"
// @Failure 500 	{object} 	web.response 	"internal server error"
// @Router /buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyers, err := b.buyerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error()) //500
			return
		}
		if len(buyers) == 0 {
			web.Success(c, http.StatusNoContent, "buyer list empty") //204
			return
		}
		web.Success(c, http.StatusOK, buyers)
	}
}

// CreateBuyer godoc
// @Summary 		Create buyer
// @Tags 			Buyers
// @Description 	post new buyer
// @Accept  		json
// @Produce  		json
// @Param 			token header int true "token"
// @Param 			domain.Buyer 	body 	domain.Buyer 	true "Buyer to create"
// @Success 201 	{object} web.response		 "buyer successfully created"
// @Failure 409     {object} web.response 		 "The card number ID provided is already registered"
// @Failure 422     {object} web.response 		 "error: ¡incomplete fields!"
// @Failure 500	    {object} web.response 		 "error while saving"
// @Router /buyers [post]
func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Buyer
		if err := c.Bind(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())//422 - Esta bien ese codigo ?
			return
		}
		if req.CardNumberID == "" || req.FirstName == "" || req.LastName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error: ¡incomplete fields!") //422
			return
		}
		//validation: buyer_code must be unique
		exist := b.buyerService.Exists(c, req.CardNumberID)
		if exist {
			web.Error(c, http.StatusConflict, "error: the buyer already exist") //409
			return
		}
		//the product doesn´t exist, let´s save the new one.
		lastId, err := b.buyerService.Save(c, req)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error: error while saving")//500
			return
		}
		req.ID = lastId

		web.Success(c, http.StatusCreated, req)
	}
}

// BuyerUpdateById godoc
// @Summary 		Update buyer by id
// @Tags 			Buyers
// @Description 	update buyer by id
// @Accept  		json
// @Produce  		json
// @Param 			token header int true "token"
// @Param 			id path int true "id"
// @Param 			domain.Buyer 	body 	domain.Buyer 	true "Buyer to update"
// @Success 200 	{object} web.response
// @Failure 400     {object} web.response	 	"bad request"
// @Failure 404     {object} web.response 		"buyer not found"
// @Failure 422     {object} web.response 		"unprocessable entity"
// @Router /buyers/{id} [patch]
func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Buyer
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())//422
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error()) //400 - Revisar codigo
			return
		}

		req.ID = id

		req, err = b.buyerService.Update(c, req)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error()) //404
			return
		}

		web.Success(c, http.StatusOK, req)
	}
}

// DeleteBuyerById godoc
// @Summary 		Delete buyer by id
// @Tags 			Buyers
// @Description 	delete buyer by id
// @Accept  		json
// @Produce  		json
// @Param 			token header int true "token"
// @Param 			id path string true "id"
// @Success 204 	{object} web.response
// @Failure 404     {object} web.response 		 "buyer not found"
// @Router /buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())//404 - revisar
			return
		}

		err = b.buyerService.Delete(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())//404
			return
		}

		web.Success(c, http.StatusNoContent, fmt.Sprintf("buyer %d deleted ", id))

	}
}