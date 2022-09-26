package main

import "fmt"

var temp int
var hum int
var pre float32

func main() {
  temp = 17
  hum = 47
  pre = 1024.5

  fmt.Printf("Temperatura: %v\n", temp)
  fmt.Printf("Humedad: %v\n", hum)
  fmt.Printf("Presión: %v\n", pre)
}
