package main

/*

Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin. 

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario: 
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.

*/

import "fmt"

func main() {
    var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

    fmt.Printf("Benjamin tiene %d años.\n", employees["Benjamin"])

    howMany := 0
    for _, age := range employees {
      if age >= 21 {
        howMany++
      }
    }

    switch howMany {
    case 0:
      fmt.Println("No hay empleados mayores de edad.")
    case 1:
      fmt.Println("Hay un empleado mayor de edad.")
    default:
      fmt.Printf("Hay %d empleados mayores de edad.\n", howMany)
    }

    employees["Federico"] = 25
    delete(employees, "Pedro")

    fmt.Println("Control: ", employees)
}
