package main

/*

Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que contenga una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

*/

import "fmt"

func main() {
  months := []string{"Diciembre", "Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Setiembre", "Octubre", "Noviembre" }

  monthNumber := 8

  fmt.Println(months[monthNumber % 12])
}
