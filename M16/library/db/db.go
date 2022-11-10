package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DataBase *sql.DB

func Init() {
	var err error
	if err = godotenv.Load(); err != nil {
		panic(err)
	}
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USERNAME"),
		os.Getenv("PASS"),
		os.Getenv("SERVER"),
		os.Getenv("PORT"),
		os.Getenv("DB"),
	)

	DataBase, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}

	if err = DataBase.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("DataBase initialized")
}
