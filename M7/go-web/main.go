package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Product struct {
    Id int
    Name string
    Colour string
    Price float64
    Stock int
    Code string
    Published bool
    CreatedAt string
}

func main() {

    router := gin.Default()

    router.GET("/hello-world", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, Mariano!",
        })
    })

    router.GET("/products", GetAll)

    router.Run()
}

func GetAll(c *gin.Context) {
    var data []byte
    var err error

    if data, err = ioutil.ReadFile("./products.json"); err != nil {
        c.JSON(500, gin.H{"error": "Oops! I fucked up!"})
    }

    var payload map[string]interface{}
    err = json.Unmarshal(data, &payload)

    if err != nil {
        c.JSON(500, gin.H{"fuck!": err})
    }
    
    c.JSON(200, payload)
}
