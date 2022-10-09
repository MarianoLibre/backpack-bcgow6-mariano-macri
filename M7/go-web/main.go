package main

import (
	"github.com/gin-gonic/gin"
)

var products []Product


func main() {
    
    FillJson()

    router := gin.Default()

    router.POST("/products", func(c *gin.Context) {
        token := c.GetHeader("token")
        if token != "un token seguro" {
            c.JSON(401, gin.H{"error": "No tiene permisos para realizar la petición solicitada."})
            return
        }

        var product Product
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        switch {
        case product.Name == "":
            c.JSON(400, gin.H{"error": "'Name' is required"})
            return
        case product.Code == "":
            c.JSON(400, gin.H{"error": "'Code' is required"})
            return
        case product.Colour == "":
            c.JSON(400, gin.H{"error": ".Colour' is required"})
            return
        case product.CreatedAt == "":
            c.JSON(400, gin.H{"error": ".CreatedAt' is required"})
            return
        case product.Price == 0.0:
            c.JSON(400, gin.H{"error": "'Price' is required"})
            return
        case product.Stock == 0:
            c.JSON(400, gin.H{"error": "'Stock' is required"})
            return
        }

        var id int
        if len(products) > 0 {
            id = products[len(products) - 1].Id + 1
        } else {
            id = 1
        }
        product.Id = id
        products = append(products, product)
        c.JSON(200, product)
    })

    router.GET("/hello-world", Hello)

    router.GET("/products", GetAll)

    misc := router.Group("/misc")
    {
        misc.GET("/", ShowContext)
        misc.GET("/hi", Hi)
    }

    router.GET("/params/:poo/:pee", PrintParam)

    router.GET("/showqueries", ShowQueries)

    router.GET("/filterby", Filter)

    router.GET("/byid/:id", FindById)

    router.Run(":8080")
}
