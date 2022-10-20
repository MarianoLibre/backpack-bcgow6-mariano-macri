package routes

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/section"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/cmd/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/buyer"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/warehouse"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/seller"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

// Middleware
func (r *router) tokenMiddleware() {
	r.rg.Use(middleware.TokenAuthMiddleware())
}

func (r *router) buildSellerRoutes() {
	// Example
	// repo := seller.NewRepository(r.db)
	// service := seller.NewService(repo)
	// handler := handler.NewSeller(service)
	// r.r.GET("/seller", handler.GetAll)
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)
	r.rg.Use(middleware.TokenAuthMiddleware())
	r.rg.POST("/sellers", handler.Create())
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.DELETE("sellers/:id", handler.Delete())
	r.rg.PATCH("sellers/:id", handler.Update())
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	r.rg.Use(middleware.TokenAuthMiddleware())
	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:id", handler.Get())
	r.rg.POST("/products", handler.Create())
	r.rg.PATCH("/products/:id", handler.Update())
	r.rg.DELETE("/products/:id", handler.Delete())
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := handler.NewSection(service)
	r.tokenMiddleware()
	r.rg.POST("/sections", handler.Create())
	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.DELETE("sections/:id", handler.Delete())
	r.rg.PATCH("sections/:id", handler.Update())
}

func (r *router) buildWarehouseRoutes() {
	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	handler := handler.NewWarehouse(service)
	// middleware
	r.rg.Use(middleware.TokenAuthMiddleware())

	r.rg.GET("/warehouses", handler.GetAll())
	r.rg.GET("/warehouses/:id", handler.Get())
	r.rg.POST("/warehouses", handler.Create())
	r.rg.DELETE("/warehouses/:id", handler.Delete())
	r.rg.PATCH("/warehouses/:id", handler.Update())
}

func (r *router) buildEmployeeRoutes() {
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	handler := handler.NewEmployee(service)

	endpoint := r.rg.Group("/employees")
	endpoint.Use(middleware.TokenAuthMiddleware())
	endpoint.POST("", handler.Create())
	endpoint.GET("", handler.GetAll())
	endpoint.GET("/:id", handler.Get())
	endpoint.PATCH("/:id", handler.Update())
	endpoint.DELETE("/:id", handler.Delete())
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)

	bg := r.rg.Group("/buyers")
	bg.Use(middleware.TokenAuthMiddleware())
	{
		bg.GET("/:id", handler.Get())
		bg.PATCH("/:id", handler.Update())
		bg.POST("/", handler.Create())
		bg.GET("/", handler.GetAll())
		bg.DELETE("/:id", handler.Delete())

	}

}
