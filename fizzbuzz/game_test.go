package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/imhotepio/letsgo_labs/fizzbuzz"
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
		assert.Equal(t, v, fizzbuzz.Compute(k))
	}
}

// Returns `FizzBuzz if number is div by 3 and 5
func ExampleCompute_fizzbuzz() {
	fmt.Println(fizzbuzz.Compute(15))
	// Output:
	// FizzBuzz
}

// Returns `Fizz if number is div by 3
func ExampleCompute_fizz() {
	fmt.Println(fizzbuzz.Compute(3))
	// Output:
	// Fizz
}

// Returns `Buzz if number is div by 5
func ExampleCompute_buzz() {
	fmt.Println(fizzbuzz.Compute(5))
	// Output:
	// Buzz
}
