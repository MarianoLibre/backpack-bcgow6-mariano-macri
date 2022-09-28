package main
/*
Bonus Track -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo.
1) Desarrolla las funciones necesarias para permitir a la empresa calcular:
    a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
        - La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
        - Dicha función deberá retornar más de un valor (salario calculado y error).
        - En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
        - En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
    b) Calcular el medio aguinaldo correspondiente al trabajador 
        - Fórmula de cálculo de aguinaldo: 
            [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
        - La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

2) Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

import (
    "fmt"
    "errors"
)

func calcSalary(hours int, value float64) (float64, error) {
    if hours < 80 {
        return 0.0, errors.New("error: el trabajador no puede haber trabajado menos de 80 horas mensuales")
    }

    total := float64(hours) * value
    if total > 100000 {
        total *= 0.9
    }
    return total, nil 
}

func calcAguinaldo (salaries []float64, months int) (float64, error) {
    var max float64
    for _, s := range salaries {
        if s < 0.0 {
            return 0.0, errors.New("error: numero negativo")
        }
        if s > max {
            max = s
        }
    }
    fmt.Printf("Mejor salario: %.2f\nMeses trabajados: %d\n", max, months)
    return max / 12.0 * float64(months), nil
}

type Employee struct {
    Name string
    Salaries []float64
}

func main() {
    employee := Employee{Name: "Mariano"}
    hoursPerMonth := []int{79, 160, 150, 40, 200}

    for _, h := range hoursPerMonth {
        salary, err := calcSalary(h, 500.0)
        if err != nil {
            fmt.Println(err)
        } else {
            employee.Salaries = append(employee.Salaries, salary)
            fmt.Printf("Sarario del mes: %.2f\n", salary)
        }
    }

    fmt.Println("Number of salaries: ", len(employee.Salaries))
    fmt.Println("Calculando aguinaldo: ")

    workingMonths := 0 
    totalMonths := len(employee.Salaries)
    switch {
    case totalMonths > 0 && totalMonths & 6 == 0:
        workingMonths = 6
    default:
        workingMonths = totalMonths % 6
    }
    aguinaldo, err := calcAguinaldo(employee.Salaries, workingMonths)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("\t%.2f\n", aguinaldo)
}
