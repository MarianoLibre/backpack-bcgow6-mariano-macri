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

    t := service.Ticket{
        Id: 666,
        Names: "Judas Priest",
        Email: "judas@priest.com",
        Destination: "London",
        Date: "11:11",
        Price: 1_000,
    }

    err2 := f.Write(&t)
    if err2 != nil {
        fmt.Println(err2)
    }

	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
