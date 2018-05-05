package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	uc := map[int]string{
		1:  "1",
		3:  fizz,
		4:  "4",
		5:  buzz,
		15: fizzbuzz,
	}

	for k, v := range uc {
		assert.Equal(t, v, compute(k))
	}
}
