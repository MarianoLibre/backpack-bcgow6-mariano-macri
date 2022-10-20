package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/pkg/web"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("token not found in environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			web.Error(c, http.StatusUnauthorized, "empty token")
			c.Abort()
			return
		}

		if token != requiredToken {
			web.Error(c, http.StatusUnauthorized, "invalid token")

			c.Abort()
			return
		}

		c.Next()
	}

}
