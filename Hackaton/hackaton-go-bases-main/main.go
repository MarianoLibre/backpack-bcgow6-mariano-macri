package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
    "github.com/bootcamp-go/hackaton-go-bases/internal/file"
)

func main() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("Oops! Something really bad happened...")
            fmt.Printf("\033[33;03m%s \033[0m\n", err)
            fmt.Println("Bye bye!")
        }
    }()

	var tickets []service.Ticket

	// Funcion para obtener tickets del archivo csv
    f := file.File{Path: "/Users/MMACRI/Bootcamp/backpack-bcgow6-mariano-macri/Hackaton/hackaton-go-bases-main/tickets.csv"} 

    data, err := f.Read()
    if err != nil {
        panic(err)
    }

    tickets = append(tickets, data...)
    
    fmt.Println("Loading tickets...")
    fmt.Println("\t", tickets[0], "\t...\n\t...\n\t...")

    bookings := service.NewBookings(tickets)

    fmt.Printf("There are %d flies booked.\n", len(tickets))
    fmt.Println("........")

    newTicket := service.Ticket{
        Id: 1666,
        Names: "Judas Priest",
        Email: "judas@priest.com",
        Destination: "London",
        Date: "11:11",
        Price: 1_000,
    }

    _, err = bookings.Create(newTicket)
    if err != nil {
        panic(err)
    }
    fmt.Println("Your new ticket has been added!")
    fmt.Println("........")

    /*
    err = f.Write(&newTicket)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Changes have been saved!")
    fmt.Println("........")
    */
    
    ticket, err := bookings.Read(1000)
    if err != nil {
        panic(err)
    }

    fmt.Println("Ticket read: ", ticket)
    fmt.Println("........")

    ticket.Price = 1_000_000
    _, err = bookings.Update(ticket.Id, ticket)
    if err != nil {
        panic(err)
    }

    fmt.Println("Ticket updated: ", ticket)
    fmt.Println("........")

    toDelete := 1
    _, err = bookings.Delete(toDelete)
    if err != nil {
        panic(err)
    }

    fmt.Println("Ticket deleted: ", toDelete)
    fmt.Println("........")

}
