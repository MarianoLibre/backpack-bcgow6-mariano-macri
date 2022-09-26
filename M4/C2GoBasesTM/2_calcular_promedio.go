package main

/*

Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

import "errors"

func main() {

}

func getAverage(grades ...int) (float64, error) {
  average := 0
  for _, grade := range grades {
    if grade < 0 {
      return 0, errors.New("Valor negativo")
    }
    average += grade
  }
  return float64(average) / float64(len(grades)), nil
}
