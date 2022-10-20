package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(w product.Service) *Product {
	return &Product{
		productService: w,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response "Products"
// @Failure 404 {object} web.response "Product not found"
// @Failure 500 {object} web.response "Failure executing get service"
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.productService.GetAll(ctx)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Failure executing get service")
			return
		}

		if len(products) == 0 {
			web.Error(ctx, http.StatusNotFound, "No products found.")
			return
		}
		web.Success(ctx, http.StatusOK, products)
	}
}

// GetProduct godoc
// @Summary Get product
// @Tags Products
// @Description get product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response "Product"
// @Failure 404 {object} web.response "Product not found"
// @Router /products/:id [get]
func (p *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Malformed ID!")
			return
		}
		product, err := p.productService.Get(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Product not found.")
			return
		}

		web.Success(ctx, http.StatusOK, product)
	}
}

// CreateProducts godoc
// @Summary Create products
// @Tags Products
// @Description create products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "Product to create"
// @Success 201 {object} web.response
// @Failure 422 {object} web.response "Gin validator error"
// @Failure 500 {object} web.response "Failure executing create service"
// @Router /api/v1/products [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.Product
		if err := ctx.Bind(&req); err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "Gin validator error")
			return
		}

		if req.Description == "" || req.ExpirationRate == 0 || req.FreezingRate == 0 || req.Height == 0 || req.Length == 0 || req.Netweight == 0 || req.ProductCode == "" || req.RecomFreezTemp == 0 || req.Width == 0 || req.ProductTypeID == 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "Incompleted fields!")
			return
		}

		//validation: product_code must be unique
		exist := p.productService.Exists(ctx, req.ProductCode)
		if exist {
			web.Error(ctx, http.StatusInternalServerError, "The card product code provided is already registered")
			return
		}

		//the product doesn´t exist, let´s save the new one.
		lastId, err := p.productService.Save(ctx, req)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "Failure executing create service")
			return
		}
		req.ID = lastId

		web.Success(ctx, http.StatusCreated, req)
	}
}

// UpdateProducts godoc
// @Summary Update products
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param user body request true "Product to update"
// @Success 200 {object} web.response "Product"
// @Failure 400 {object} web.response "Malformed ID"
// @Failure 404 {object} web.response "Product not found"
// @Failure 422 {object} web.response "Gin validator error"
// @Router /api/v1/products/:id [patch]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Malformed ID")
			return
		}

		var prodReq domain.Product
		if errReq := ctx.Bind(&prodReq); err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, errReq.Error())
			return
		}

		//validation: product_code must exist for updating
		exist := p.productService.Exists(ctx, prodReq.ProductCode)
		if !exist {
			web.Error(ctx, http.StatusNotFound, "The product specified doesn't exist in database.")
			return
		}

		productOriginal, err := p.productService.Get(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "Product not found.")
			return
		}

		productOriginal = BuildStruct(&productOriginal, &prodReq)

		errUpd := p.productService.Update(ctx, productOriginal)
		if errUpd != nil {
			web.Error(ctx, http.StatusBadRequest, errUpd.Error())
			return
		}

		web.Success(ctx, http.StatusOK, productOriginal)
	}
}

// UpdateProducts godoc
// @Summary Delete product
// @Tags Products
// @Description delete product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response "Product"
// @Failure 400 {object} web.response "Malformed ID"
// @Failure 404 {object} web.response "Product not found"
// @Router /api/v1/products/:id [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "Malformed ID")
			return
		}

		errDelete := p.productService.Delete(ctx, id)
		if errDelete != nil {
			web.Error(ctx, http.StatusNotFound, "Product not found")
			return
		}

		web.Success(ctx, http.StatusNoContent, "") //As http.StatusNoContent has no response, any response message will be ignored.
	}
}

func BuildStruct(prodOriginal *domain.Product, prodReq *domain.Product) domain.Product {
	if prodReq.Description != "" {
		prodOriginal.Description = prodReq.Description
	}
	if prodReq.ExpirationRate != 0 {
		prodOriginal.ExpirationRate = prodReq.ExpirationRate
	}
	if prodReq.FreezingRate != 0 {
		prodOriginal.FreezingRate = prodReq.FreezingRate
	}
	if prodReq.Height != 0 {
		prodOriginal.Height = prodReq.Height
	}
	if prodReq.Length != 0 {
		prodOriginal.Length = prodReq.Length
	}
	if prodReq.Netweight != 0 {
		prodOriginal.Netweight = prodReq.Netweight
	}
	if prodReq.ProductCode != "" {
		prodOriginal.ProductCode = prodReq.ProductCode
	}
	if prodReq.RecomFreezTemp != 0 {
		prodOriginal.RecomFreezTemp = prodReq.RecomFreezTemp
	}
	if prodReq.Width != 0 {
		prodOriginal.Width = prodReq.Width
	}
	if prodReq.ProductTypeID != 0 {
		prodOriginal.ProductTypeID = prodReq.ProductTypeID
	}
	return *prodOriginal
}
