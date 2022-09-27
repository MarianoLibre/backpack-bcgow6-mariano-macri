package main

/*

Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

import (
	"fmt"
	"os"
)

type Product struct {
    Id int
    Price float64
    Quantity int
}

type Stock struct {
    Items []Product
}

func main() {
    
    p1 := Product{1, 10.5, 1}
    p2 := Product{2, 30.5, 6}
    p3 := Product{3, 100.5, 10}
    p4 := Product{4, 1.6, 301}

    stock := Stock{[]Product{p1, p2, p3, p4}}

    info := "Id;Price;Quantity\n"

    for _, p := range stock.Items {
        newLine := fmt.Sprintf("%d;%.2f;%d\n", p.Id, p.Price, p.Quantity)
        info += newLine
    }

    err := os.WriteFile("./sample.csv", []byte(info), 0644)
    if err != nil {
        fmt.Println("Oops!")
    }
    fmt.Println("Yeah!")
}
