package main

import "github.com/gin-gonic/gin"

 
func ShowContext(c *gin.Context) {
    payload := make(map[string]interface{})

    payload["Body"] = c.Request.Body
    payload["Method"] = c.Request.Method
    payload["Header"] = c.Request.Header

    c.JSON(200, payload)
}

func Hello(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, Mariano!",
    })
}

func Hi(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hi, Mariano!",
    })
}

func PrintParam(c *gin.Context) {
    c.String(200, "Params: %v", c.Params)
}
