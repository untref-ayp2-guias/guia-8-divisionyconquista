package guia8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjercicio1(t *testing.T) {
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 1))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 2))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 3))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 4))
	assert.True(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, BusquedaBinaria([]int{1, 2, 3, 4, 5}, 6))
}

func TestEjercicio2(t *testing.T) {
	assert.Equal(t, 8, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
	assert.Equal(t, 0, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1))
	assert.Equal(t, 4, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	assert.Equal(t, -1, BusquedaTernariaRecursiva([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10))
}

func TestEjercicio3(t *testing.T) {
	assert.Equal(t, 4, Pico([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}))
	assert.Equal(t, 8, Pico([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, 0, Pico([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestEjercicio4(t *testing.T) {
	assert.Equal(t, 0, Multiplicacion(0, 6))
	assert.Equal(t, 0, Multiplicacion(4, 0))
	assert.Equal(t, 40, Multiplicacion(5, 8))
	assert.Equal(t, 27, Multiplicacion(9, 3))
}

func TestEjercicio5(t *testing.T) {
	assert.Equal(t, 4, ElementoEnSuPosicion([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 1, ElementoEnSuPosicion([]int{0, 1, 2, 3, 5, 6, 7, 8, 9}))
	assert.Equal(t, 6, ElementoEnSuPosicion([]int{-9, -4, -3, 1, 2, 5, 6, 7, 8}))
}
