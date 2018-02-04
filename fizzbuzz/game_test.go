package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	uc := map[int]string{
		0:  "FizzBuzz",
		1:  "1",
		3:  "Fizz",
		4:  "4",
		5:  "Buzz",
		15: "FizzBuzz",
	}

	for k, v := range uc {
		assert.Equal(t, v, compute(k))
	}
}
