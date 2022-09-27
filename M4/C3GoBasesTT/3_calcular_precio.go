package main
/*

Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

import (
    "fmt"
)

// ------------------------------------------------------- //
type Product struct {
    Name string
    Price float64
    Quantity int
}

func newProduct(name string, price float64, quantity int) Product {
    return Product{name, price, quantity}
}

// ------------------------------------------------------- //
type Service struct {
    Name string
    Price float64
    WorkMinutes int
}

func newService(name string, price float64) Service {
    return Service{Name: name, Price: price}
}

// ------------------------------------------------------- //
type Maintenace struct {
    Name string
    Price float64
}

func newMaintenance(name string, price float64) Maintenace {
    return Maintenace{Name: name, Price: price}
}

// ------------------------------------------------------- //
func sumProducts(products []Product, c chan float64) {
    var total float64
    for _, product := range products {
        total += product.Price * float64(product.Quantity)
    }
    c <- total
}

// ------------------------------------------------------- //
func sumServices(services []Service, c chan float64) {
    var total float64
    for _, service := range services {
        total += service.Price * float64((service.WorkMinutes / 30 + service.WorkMinutes % 30))
    }
    c <- total
}

// ------------------------------------------------------- //
func sumMaintenance(maintenances []Maintenace, c chan float64) {
    var total float64
    for _, maintenance := range maintenances {
        total += maintenance.Price
    }
    c <- total
}

// ------------------------------------------------------- //
func main() {
    products := []Product{
        {"Celular", 10.0, 1},
        {"Cafe", 10.0, 1},
        {"Tablet", 10.0, 1},
    }
    
    services := []Service{
        {"Limpieza", 10.0, 31},
        {"Envios", 10.0, 1},
    }

    maintenances := []Maintenace{
        {"Cambio de modulo", 10.0},
        {"Bateria", 10.0},
    }

    c := make(chan float64)

    go sumProducts(products, c)
    totalProductos := <- c

    go sumMaintenance(maintenances, c)
    totalMaintenance := <- c

    go sumServices(services, c)
    totalServices := <- c
    fmt.Printf("%-20s| %15.2f\n", "Productos", totalProductos)
    fmt.Printf("%-20s| %15.2f\n", "Mantenimiento", totalMaintenance)
    fmt.Printf("%-20s| %15.2f\n", "Servicios", totalServices)
    for i := 0; i < 37; i++ {
        if i == 20 {
            fmt.Print("+")
        } else {
            fmt.Print("-")
        }
    }
    fmt.Println()
    fmt.Printf("%-20s| %15.2f\n", "Total", totalMaintenance + totalProductos + totalServices)
}
