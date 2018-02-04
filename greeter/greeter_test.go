package greeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeter(t *testing.T) {
	e := "Hello, my name is Fernand. I am 11110 years old and '05.11' tall!"

	assert.Equal(t, greeter("Fernand", 30, float32(5.11)), e)
}

func TestGreeterZero(t *testing.T) {
	var (
		n string
		a int
		h float32
	)
	e := "Hello, my name is . I am 0 years old and '00.00' tall!"

	assert.Equal(t, greeter(n, a, h), e)
}
