package main

/*

Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior
Ejemplo:

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)
 
...
 
minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)
 
...
 
minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

*/

import "errors"

func main() {

}

func minimum(values ...int) (float64, error) {
  result := values[0]

  for _, v := range values {
    if v < 0 {
      return 0, errors.New("Valor negativo")
    }
    if v < result {
      result = v
    }
  }

  return float64(result), nil
}

func maximum(values ...int) (float64, error) {
  result := values[0]

  for _, v := range values {
    if v < 0 {
      return 0, errors.New("Valor negativo")
    }
    if v > result {
      result = v 
    }
  }

  return float64(result), nil
}

func average(values ...int) (float64, error) {
  result := 0

  for _, v := range values {
    if v < 0 {
      return 0, errors.New("Valor negativo")
    }
    result += v
  }

  return float64(result) / float64(len(values)), nil
}

    
func operation(op string) (func(values ...int) (float64, error), error) {
  switch op {
  case "minimum":
    return minimum, nil
  case "maximum":
    return maximum, nil
  case "average":
    return average, nil
  default:
    return nil, errors.New("Invalid argument")
  }
}
