package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
    num1 := 8
    num2 := 3
    resultadoEsperado := 5

    resultado := Restar(num1, num2)

    assert.Equal(t, resultadoEsperado, resultado, "Are you kidding' me!?")

}

func TestDividir(t *testing.T) {
	num1, num2 := 100, 2
	resultadoEsperado := 50

	resultado, err := Dividir(num1, num2)
	assert.Equal(t, resultado, resultadoEsperado, "Ajaja@")
	assert.Nil(t, err, "Fiuuu!")

	num2 = 0
	resultadoEsperado = 0

	resultado, err = Dividir(num1, num2)
	assert.Equal(t, resultado, resultadoEsperado, "Really?")
	assert.NotNil(t, err, "Roger that!")
}
