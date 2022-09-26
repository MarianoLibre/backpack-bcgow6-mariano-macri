package main

/*

Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10%.
*/


func main() {

}

func tax(salary float64) float64 {
  switch {
  case salary > 150000:
    return salary * 0.27
  case salary > 50000:
    return salary * 0.17
  default:
    return 0
  }
}
