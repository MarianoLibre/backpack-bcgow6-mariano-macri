package main

import (

	"github.com/gin-gonic/gin"
)

func main() {
    
    FillJson()

    router := gin.Default()

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

    router.Run(":8080")
}
