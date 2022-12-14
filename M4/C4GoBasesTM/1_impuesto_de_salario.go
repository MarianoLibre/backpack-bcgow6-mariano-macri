package main
/*

Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

import (
    "fmt"
)

type myError struct {
    msg string 
}

func (e *myError) Error() string {
    return e.msg
}

func checkSalary(salary int) (string, error) {
    msg := "Debe pagar impuesto"
    if salary < 150000 {
        return "", &myError {
            "error: el salario ingresado no alcanza el mínimo imponible",
        }
    }
    return msg, nil
}

func main() {
    salary := 200000
    msg, err := checkSalary(salary)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(msg)
}
