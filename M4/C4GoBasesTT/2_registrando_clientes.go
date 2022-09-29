package main

/*
Ejercicio 2 - Registrando clientes

El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio
*/

import (
	"errors"
	"fmt"
	"os"
)

type Client struct {
    File int
    FirstName string
    LastName string
    Id int
    PhoneNunber int
    Address string
}

type NameError struct {
    msg string
}

func (ne *NameError) Error() string {
    return ne.msg
}

func idGenerator() func(string) (int, error) {
    id := 0
    return func(name string) (int, error) {
        if len(name) < 2 {
            return 0, &NameError{"Name is too short"}
        }
        id++
        return id, nil
    }
}

func openFile(name string) {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
    }()
        
    file, err := os.ReadFile(name)
    if err != nil {
        panic("el archivo indicado no fue encontrado o está dañado")
    }
    if len(file) > 0 {
        fmt.Println("El archivo existe.")
    }
}

func validateClient(client Client) (bool, error) {
    switch {
    case client.Address == "":
        return false, errors.New("error: Address can not be empty")
    case client.File == 0:
        return false, errors.New("error: File can not be empty")
    case client.FirstName == "":
        return false, errors.New("error: First name can not be empty")
    case client.LastName == "":
        return false, errors.New("error: Last name can not be empty")
    case client.Id == 0:
        return false, errors.New("error: Id can not be empty")
    case client.PhoneNunber == 0:
        return false, errors.New("error: PhoneNunber can not be empty")
    }
    return true, nil
}

func main() {
    defer func() {
        err := recover()
        fmt.Println("Fin de la ejecución")
        if err != nil {
            fmt.Println("Se detectaron varios errores en tiempo de ejecución")
            fmt.Println("No han quedado archivos abiertos")
        }
    }()

    genId := idGenerator()
        
    mariano := new(Client)

    mariano.FirstName = "Mariano"
    id, err1 := genId(mariano.FirstName)

    if err1 != nil {
        panic(err1)
    }
    
    openFile("./customers.txt.")
    fmt.Println("Mariano's id is ", id)
    
    isOk, err2 := validateClient(*mariano)
    if err2 != nil {
        panic("Fields are missing")
    }
    if isOk {
        fmt.Println("Registering client")
    }
}
