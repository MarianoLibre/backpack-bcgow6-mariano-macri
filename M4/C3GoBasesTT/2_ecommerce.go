package main
/*

Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

import (
    "fmt"
)

type Product struct {
    Name string
    Price float64
    Quantity int
}

type User struct {
    FirstName string
    LastName string
    Email string
    Products []Product
}

func newProduct(name string, price float64) Product {
    return Product{Name: name, Price: price}
}

func (user *User) addProduct(product *Product, quantity int) {
    product.Quantity = quantity
    user.Products = append(user.Products, *product)
}

func (user *User) deleteProducts() {
    user.Products = []Product{}
}

func main() {
    user := User{
        FirstName: "Mariano",
        LastName: "Libre",
        Email: "mimail@mercadolibre.com",
    }

    celu := newProduct("celu", 20000.5)

    fmt.Println(celu)
    fmt.Println(user)
    user.addProduct(&celu, 1)
    fmt.Println(celu)
    fmt.Println(user)
    
    payanas := newProduct("Payanas", 100.4)

    fmt.Println(payanas)
    fmt.Println(user)
    user.addProduct(&payanas, 5)
    fmt.Println(payanas)
    fmt.Println(user)
    
}
