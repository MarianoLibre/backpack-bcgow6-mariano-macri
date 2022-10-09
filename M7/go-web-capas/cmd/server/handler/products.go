package handler


import (
   "github.com/gin-gonic/gin"
   "github.com/MarianoLibre/backpack-bcgow6-mariano-macri/M7/go-web-capas/internal"
)

type request struct {
    Name        string      `json:"name"`
    Colour      string      `json:"colour"`
    Price       float64     `json:"price"`
    Stock       int         `json:"stock"`
    Code        string      `json:"code"`
    Published   bool        `json:"published"`
    CreatedAt   string      `json:"created-at"`
}

type Product struct {
   service products.Service
}

func NewProduct(p products.Service) *Product {
   return &Product{
       service: p,
   }
}

func (c *Product) GetAll() gin.HandlerFunc {
   return func(ctx *gin.Context) {
       token := ctx.Request.Header.Get("token")
       if token != "123456" {
          ctx.JSON(401, gin.H{
             "error": "token inválido",
          })
          return
       }
  
       p, err := c.service.GetAll()
       if err != nil {
          ctx.JSON(404, gin.H{
             "error": err.Error(),
          })
          return
       }
       ctx.JSON(200, p)
    }
  }

func (c *Product) Store() gin.HandlerFunc {
   return func(ctx *gin.Context) {
       token := ctx.Request.Header.Get("token")
       if token != "123456" {
          ctx.JSON(401, gin.H{ "error": "token inválido" })
          return
       }
       var req request
       if err := ctx.Bind(&req); err != nil {
          ctx.JSON(404, gin.H{
             "error": err.Error(),
          })
          return
       }
       p, err := c.service.Store(req.Name, req.Colour, req.Code, req.CreatedAt, req.Stock, req.Price, req.Published)
       if err != nil {
          ctx.JSON(404, gin.H{ "error": err.Error() })
          return
       }
       ctx.JSON(200, p)
    }
  }
