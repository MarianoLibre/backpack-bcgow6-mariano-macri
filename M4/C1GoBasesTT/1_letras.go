package main

/*
Ejercicio 1 - Letras de una palabra

La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/
import "fmt"


func main() {
  word := "Palabra"

  fmt.Printf("La palabra tiene %d letras.\n", len(word))
  
  for i, letter := range word {
    fmt.Printf("La letra numero %d es %q\n", i + 1, letter)
  }
}
