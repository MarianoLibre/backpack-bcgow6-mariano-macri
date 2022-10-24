package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MarianoLibre/go-web-capas/cmd/server/handler"
	"github.com/MarianoLibre/go-web-capas/internal/products"
	"github.com/MarianoLibre/go-web-capas/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.PATCH("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_UpdateProduct(t *testing.T) {
    r := createServer()

	// Creata a product
	req, rr := createRequestTest(http.MethodPost, "/products/",
	`{
		"name": "new product",
		"colour": "new colour",
		"price": 123.45,
		"stock": 100,
		"code": "new code",
		"published": true,
		"created-at": "new date"
	}`)

    r.ServeHTTP(rr, req)
    assert.Equal(t, 200, rr.Code)

	req, rr = createRequestTest(http.MethodPatch, "/products/1",
	`{
		"name": "updated product",
		"colour": "updated colour",
		"price": 123.45,
		"stock": 100,
		"code": "updated code",
		"published": true,
		"created-at": "updated date"
	}`)

    r.ServeHTTP(rr, req)
    assert.Equal(t, 200, rr.Code)

	req, rr = createRequestTest(http.MethodDelete, "/products/1", "")

    r.ServeHTTP(rr, req)

    assert.Equal(t, 200, rr.Code)
}

