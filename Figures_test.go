package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {

	assert := assert.New(t)

	var s Figure
	var c Figure

	for j := 1; j <= 10; j++ {
		s = Square{a: float64(j)}
		c = Circle{r: float64(j)}

		assert.Equal(s.Area(), float64(j*j), "Error in the calculation the area of the square.")
		assert.Equal(s.Perimeter(), 4*float64(j), "Error in the calculation the perimeter of the square.")

		assert.Equal(c.Area(), math.Pi*float64(j*j), "Error in the calculation the area of the circle.")
		assert.Equal(c.Perimeter(), 2*math.Pi*float64(j), "Error in the calculation the perimeter of the circle.")

	}
}
