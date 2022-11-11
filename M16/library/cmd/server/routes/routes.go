package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"library.com/cmd/server/handler"
	"library.com/internal/book"
	"library.com/internal/library"
	"library.com/internal/loan"
	"library.com/internal/user"
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

	r.buildBookRoutes()
	r.buildLibraryRoutes()
	r.buildUserRoutes()
	r.buildLoanRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildUserRoutes() {
	repo := user.NewRepository(r.db)
	svc := user.NewService(repo)
	h := handler.NewUser(svc)

	r.rg.GET("/users", h.GetAllUsers())
	r.rg.GET("/users/:id", h.GetUser())
	r.rg.POST("/users", h.StoreUser())
	r.rg.PUT("/users/:id", h.UpdateUser())
	r.rg.DELETE("/users/:id", h.DeleteUser())
}

func (r *router) buildBookRoutes() {
	repo := book.NewRepository(r.db)
	svc := book.NewService(repo)
	h := handler.NewBook(svc)

	r.rg.GET("/books", h.GetAllBooks())
	r.rg.GET("/books/:id", h.GetBook())
	r.rg.POST("/books", h.StoreBook())
	r.rg.PUT("/books/:id", h.UpdateBook())
	r.rg.DELETE("/books/:id", h.DeleteBook())
}

func (r *router) buildLibraryRoutes() {
	repo := library.NewRepository(r.db)
	svc := library.NewService(repo)
	h := handler.NewLibrary(svc)

	r.rg.GET("/libraries", h.GetAllLibraries())
	r.rg.GET("/libraries/:id", h.GetLibrary())
	r.rg.POST("/libraries", h.StoreLibrary())
	r.rg.PUT("/libraries/:id", h.UpdateLibrary())
	r.rg.DELETE("/libraries/:id", h.DeleteLibrary())
}

func (r *router) buildLoanRoutes() {
	repo := loan.NewRepository(r.db)
	svc := loan.NewService(repo)
	h := handler.NewLoan(svc)

	r.rg.GET("/loans", h.GetAllLoans())
	r.rg.GET("/loans/:id", h.GetLoan())
	r.rg.POST("/loans", h.StoreLoan())
	r.rg.PUT("/loans/:id", h.UpdateLoan())
	r.rg.DELETE("/loans/:id", h.DeleteLoan())
}
