package main

/*
Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos: 
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.


Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos. 
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

*/

// Crear una estructura “tienda” que guarde una lista de productos. 
type product struct {
    Type string
    Name string
    Price float64
}

// Crear una interface “Producto” que tenga el método “CalcularCosto”
type Product interface {
    CalculateCost() float64
}

// Interface Producto:
// El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
func (p product) CalculateCost() float64 {
// Sus costos adicionales son:
// Pequeño: El costo del producto (sin costo adicional)
// Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
// Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.
    switch p.Type {
    case "medium":
        return p.Price * 0.03      
    case "big":
        return p.Price * 0.06 + 2500
    default:
        return 0.0
    }
}

// Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
type Ecommerce interface {
    Total() float64
    Add(product)
 }

// Crear una estructura “tienda” que guarde una lista de productos. 
type Store struct {
    Products []product
}

// - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
func (s Store) Total() float64 {
    result := 0.0
    for _, v := range s.Products {
        result += v.Price + v.CalculateCost()
    }
    return result
}

// - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
func (s *Store) Add(p product) {
    s.Products = append(s.Products, p)
}

// Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
func newProduct(_type, name string, price float64) product {
    return product{_type, name, price}
}

// Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
func newStore() Ecommerce {
    return &Store{}
}

func main() {

}
