package main

import (
	"github.com/gin-gonic/gin"
	"library.com/cmd/server/routes"
	"library.com/db"
)

func main() {
	db.Init()

	gin.SetMode(gin.ReleaseMode)

	eng := gin.Default()

	router := routes.NewRouter(eng, db.DataBase)
	router.MapRoutes()

	if err := eng.Run("localhost:8888"); err != nil {
		panic(err)
	}
}
