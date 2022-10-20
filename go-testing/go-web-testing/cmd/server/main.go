package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/docs"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/cmd/server/routes"

	_"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title MELI Bootcamp API FP
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://localhost:8080/
func main() {
	// NO MODIFICAR
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/melisprint")
	if err != nil {
		panic(err)
	}
	docs.SwaggerInfo.Host = os.Getenv("HOST")

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	eng := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}


	router := routes.NewRouter(eng, db)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}
