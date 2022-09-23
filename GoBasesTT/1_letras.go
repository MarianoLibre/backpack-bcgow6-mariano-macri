package main

import "fmt"


func main() {
  word := "Palabra"

  fmt.Printf("La palabra tiene %d letras.\n", len(word))
  
  for i, letter := range word {
    fmt.Printf("La letra numero %d es %q\n", i + 1, letter)
  }
}
