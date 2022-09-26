package main

import "fmt"

/*

Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

type Student struct {
  FistName string
  LastName string
  ID int
  Date string
}

func (s Student) details() {
  fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", s.FistName, s.LastName, s.ID, s.Date)
}

func main() {
    mmacri := Student{"Mariano", "Macri", 30000000, "30/30/3000"}

    mmacri.details()
}
