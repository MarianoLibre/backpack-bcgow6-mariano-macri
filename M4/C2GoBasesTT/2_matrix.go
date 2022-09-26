package main

import "fmt"

/*

Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
  Values []float64
  Height int
  Width int
  IsQuadratic bool
  MaxValue int
}

func (m *Matrix) Set(values ...float64) { m.Values = values }

func (m Matrix) Print() {
  for line := 0; line < m.Height; line++ {
    for value := 0; value < m.Width; value++ {
      fmt.Printf("%f\t", m.Values[line * m.Width + value])
    }
    fmt.Printf("\n")
  }
}

func main() {
    m := Matrix{Height: 3, Width: 3, IsQuadratic: true}

    m.Set(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)

    m.Print()
}
