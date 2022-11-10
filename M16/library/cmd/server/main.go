package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"library.com/db"
)

func main() {

	fmt.Println("Example")
	fmt.Println(os.Getenv("SHELL"))
	fmt.Println("pre: ", os.Getenv("SOMEVAR"))

	if godotenv.Load() != nil {
		fmt.Println("Can't load env vars")
		os.Exit(1)
	}

	fmt.Println("post: ", os.Getenv("SOMEVAR"))
	fmt.Println("post: [U]", os.Getenv("USERNAME"))
	fmt.Println("post: [u]", os.Getenv("user"))
	fmt.Println("post: ", os.Getenv("PASS"))
	fmt.Println("post: ", os.Getenv("DB"))

	fmt.Println("End")

	db.Init()
	db.DataBase.Ping()
}
