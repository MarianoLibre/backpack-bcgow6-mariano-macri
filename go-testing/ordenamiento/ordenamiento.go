package ordenamiento

import (
	"fmt"
	"time"
)


func BubbleSort(a []int) []int {
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
	return arr
}
