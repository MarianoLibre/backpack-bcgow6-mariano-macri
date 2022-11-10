package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"library.com/cmd/server/handler"
	"library.com/internal/book"
	"library.com/internal/library"
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

	r.rg.GET("/librarys", h.GetAllLibraries())
	r.rg.GET("/librarys/:id", h.GetLibrary())
	r.rg.POST("/librarys", h.StoreLibrary())
	r.rg.PUT("/librarys/:id", h.UpdateLibrary())
	r.rg.DELETE("/librarys/:id", h.DeleteLibrary())
}
