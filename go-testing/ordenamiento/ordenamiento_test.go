package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBubbleSort(t *testing.T) {
	arr := []int{4, 1, 3, 2}
	expected := []int{1, 2, 3, 4}

	output := BubbleSort(arr)

	assert.Equal(t, expected, output, "WTF!")
}
