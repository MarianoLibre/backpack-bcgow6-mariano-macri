package main

/*

Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.


Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.


ejemplo:

const (
   dog = "dog"
   cat = "cat"
)
 
...
 
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)
 
...

*/


func main() {

}

func Animal(animal string) (func(int) int, string) {
  switch animal {
  case "dog":
    return dogFood, ""
  case "cat":
    return catFood, ""
  case "hamster":
    return hamsterFood, ""
  case "tarantula":
    return tarantulaFood, ""    
  default:
    return nil, "No se encontró el animal '" + animal + "'."
  }
}

func dogFood(count int) int { return count * 10 * 1000 }

func catFood(count int) int { return count * 5 * 1000 }

func hamsterFood(count int) int { return count * 250 }

func tarantulaFood(count int) int { return count * 150 }

