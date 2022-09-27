package main
/*

Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.

*/

import (
    "fmt"
)

type User struct {
    FirstName string
    LastName string
    Age int
    Email string
    Password string
}

func (u *User) ChangeName(first, last string) {
    u.FirstName = first
    u.LastName = last
    fmt.Printf("%p\n", u)
}

func (u *User) ChangeAge(age int) {
    u.Age = age
}

func (u *User) ChangeEmail(email string) {
    u.Email = email
}

func (u *User) ChangePassword(pass string) {
    u.Password = pass
}

func main() {
    user := User{
        "Mariano",
        "Libre",
        44,
        "mimail@mail.com",
        "1234",
    }

    fmt.Println(user)
    fmt.Printf("%p\n", &user)
    
    user.ChangeName("Nano", "Free")
    user.ChangeAge(18)
    user.ChangeEmail("nuevomail@mail.com")
    user.ChangePassword("0000")

    fmt.Println(user)
}
