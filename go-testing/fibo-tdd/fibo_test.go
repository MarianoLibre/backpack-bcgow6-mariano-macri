package fibotdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases fibo
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144.
//
// F0 = 0, F1 = 1
// FN = F(N-1) + F(N-2)
func TestFibo(t *testing.T) {
	seq := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

	for i := range seq {
		expected := seq[:i]
		out := Fibo(i)

		assert.Equal(t, expected, out)
	}
}
