package benchmark

import (
	"crypto/sha1"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSum256(b *testing.B) {
	data := []byte("Digital House impulsando la transformacion digital")
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}
func BenchmarkSum(b *testing.B) {
	data := []byte("Digital House impulsando la transformacion digital")
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func TestFactorialIter(t *testing.T) {
	assert.Equal(t, FactorialIter(0), 1)
	assert.Equal(t, FactorialIter(1), 1)
	assert.Equal(t, FactorialIter(2), 2)
	assert.Equal(t, FactorialIter(3), 3 * 2)
	assert.Equal(t, FactorialIter(4), 4 * 3 * 2)
	assert.Equal(t, FactorialIter(5), 5 * 4 * 3 * 2)
	assert.Equal(t, FactorialIter(6), 6 * 5 * 4 * 3 * 2)
	assert.Equal(t, FactorialIter(7), 7 * 6 * 5 * 4 * 3 * 2)
}

func TestFactorialRecur(t *testing.T) {
	assert.Equal(t, FactorialRecur(0), 1)
	assert.Equal(t, FactorialRecur(1), 1)
	assert.Equal(t, FactorialRecur(2), 2)
	assert.Equal(t, FactorialRecur(3), 3 * 2)
	assert.Equal(t, FactorialRecur(4), 4 * 3 * 2)
	assert.Equal(t, FactorialRecur(5), 5 * 4 * 3 * 2)
	assert.Equal(t, FactorialRecur(6), 6 * 5 * 4 * 3 * 2)
	assert.Equal(t, FactorialRecur(7), 7 * 6 * 5 * 4 * 3 * 2)
}

func BenchmarkFactorialIter(b *testing.B) {
	for i := 0; i < 10_000; i++ {
		_ = FactorialIter(i)
	}
}

func BenchmarkFactorialRecur(b *testing.B) {
	for i := 0; i < 10_000; i++ {
		_ = FactorialRecur(i)
	}
}
