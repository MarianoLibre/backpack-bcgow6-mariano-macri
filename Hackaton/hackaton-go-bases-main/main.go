package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
    "github.com/bootcamp-go/hackaton-go-bases/internal/file"
)

func main() {
    f := file.File{Path: "/Users/MMACRI/Bootcamp/backpack-bcgow6-mariano-macri/Hackaton/hackaton-go-bases-main/tickets.csv"} 
    data, err := f.Read()
    if err != nil {
        panic(err)
    }
    fmt.Println(data[0])
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
