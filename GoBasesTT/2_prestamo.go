package main
/*
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000. 
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

import "fmt"

const (
  minAge = 22
  minYears = 1
  minIncome = 100000
)

func main() {
    age := 25
    isEmployed := true
    seniority := 3
    income := 100000.10

    if age <= minAge {
      fmt.Println("Tienes que ser mayor de 22 para recibir el préstamo.")
    } else if !isEmployed {
      fmt.Println("Tienes que estar empleado para recibir el préstamo.")
    } else if seniority <= minYears {
      fmt.Println("Tienes que tener más de un año de antigüedad par recibir el préstamo.")
    } else {
      fmt.Println("Cumples todos los requisitos. Te vamos a dar la plata :)")
      if income > minIncome {
        fmt.Println("Y no te vamos a cobrar interesees.")
      }
    }
}
