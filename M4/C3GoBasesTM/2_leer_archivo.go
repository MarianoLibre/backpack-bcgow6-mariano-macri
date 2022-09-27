package main

/*

Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    data, err := os.ReadFile("./sample.csv")
    if err != nil {
        fmt.Println("Oops!")
    }

    fdata := strings.Split(string(data), "\n")
    var total float64
    for _, line := range fdata {
        if len(line) > 0 {
            fline := strings.Split(line, ";")
            fmt.Printf("%s\t%10s\t%10s\n", fline[0], fline[1], fline[2])
            price, err := strconv.ParseFloat(fline[1], 64)
            q, err2 := strconv.ParseFloat(fline[2], 64)
            if err == nil && err2 == nil {
                total += price * q
            }
        }
    }
    fmt.Printf("%s\t%10.2f\n", "Total:", total)
}
