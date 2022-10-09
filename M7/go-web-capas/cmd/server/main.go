package main


import (
    "github.com/gin-gonic/gin"
    "github.com/MarianoLibre/backpack-bcgow6-mariano-macri/tree/main/M7/go-web-capas/internal/products"
)


func main() {
   repo := products.NewRepository()
   service := products.NewService(repo)
   p := handler.NewProduct(service)

   r := gin.Default()
   pr := r.Group("/products")
   pr.POST("/", p.Store())
   pr.GET("/", p.GetAll())
   r.Run()
}
