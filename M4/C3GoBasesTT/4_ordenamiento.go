package main

/*

Opcional - Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
)


func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)
}

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func insertSort(a []int, c chan []int) {
    // I make a copy so all the funcs can work on the original
    // slice without modifying it.
    start := time.Now()
    arr := make([]int, len(a))
    copy(arr, a)
    i := 1
    for i < len(arr) {
        j := i
        for j > 0 && arr[j - 1] > arr[j] {
            temp := arr[j]
            arr[j] = arr[j - 1]
            arr[j - 1] = temp
            j--
        }
        i++
    }
    end := time.Now()
    fmt.Printf("\tInsert sort took %v\n", end.Sub(start))
    c <- arr
}

func bubbleSort(a []int, c chan []int) {
    // I make a copy so all the funcs can work on the original
    // slice without modifying it.
    start := time.Now()
    arr := make([]int, len(a))
    copy(arr, a)
    n := len(arr)
    for {
        swapped := false
        for i := 1; i < n; i++ {
            if arr[i - 1] > arr[i] {
                temp := arr[i]
                arr[i] = arr[i - 1]
                arr[i - 1] = temp
                swapped = true
            }
        }
        if !swapped {
            break
        }
    }
    end := time.Now()
    fmt.Printf("\tBubble sort took %v\n", end.Sub(start))
    c <- arr
}

func selectSort(a []int, c chan []int) {
    // I make a copy so all the funcs can work on the original
    // slice without modifying it.
    start := time.Now()
    arr := make([]int, len(a))
    copy(arr, a)
    for i := 0; i < len(arr); i++ {
        min := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }
        if min != i {
            temp := arr[i]
            arr[i] = arr[min]
            arr[min] = temp
        }
    }
    end := time.Now()
    fmt.Printf("\tSelect sort took %v\n", end.Sub(start))
    c <- arr
}

func printHeader(n int) {
    fmt.Println("Array length is ", n)
}

func main() {
    arr := rand.Perm(100)
    c := make(chan []int)
    printHeader(len(arr))    
    go insertSort(arr, c)
    <- c
    go bubbleSort(arr, c)
    <- c
    go selectSort(arr, c)
    <- c

    arr = rand.Perm(1000)
    printHeader(len(arr))    
    go insertSort(arr, c)
    <- c
    go bubbleSort(arr, c)
    <- c
    go selectSort(arr, c)
    <- c

    arr = rand.Perm(10000)
    printHeader(len(arr))    
    go insertSort(arr, c)
    <- c
    go bubbleSort(arr, c)
    <- c
    go selectSort(arr, c)
    <- c
}
