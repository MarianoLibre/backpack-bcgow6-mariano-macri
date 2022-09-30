package file

import (
	"os"
	"strings"
    "strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
    data, err := os.ReadFile(f.Path)
    if err != nil {
        return nil, err
    }
    content := strings.Split(string(data), "\n")
    tickets := []service.Ticket{}
    for _, line := range content {
        fields := strings.Split(line, ",")
        id, err := strconv.Atoi(fields[0])
        if err != nil {
            return nil, err
        }
        price, err := strconv.Atoi(fields[5])
        if err != nil {
            return nil, err
        }
        newTicket := service.Ticket{
            Id: id,
            Names: fields[1],
            Email: fields[2],
            Destination: fields[3],
            Date: fields[4],
            Price: price,
        }
        tickets = append(tickets, newTicket)
    }
    
	return tickets, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}

