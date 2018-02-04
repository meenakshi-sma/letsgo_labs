// FizzBuzz game calculator
package fizzbuzz

import (
	"fmt"
)

// compute returns a FizzBuzz number based on given inputs.
// Returns `FizzBuzz` if number div by 3 and 5.
// Fizz if div by 3.
// Buzz if div by 5
// Otherwise the given number
func compute(n int) (s string) {
	switch {
	case n%3 == 0 && n%5 == 0:
		s = "FizzBuzz"
	case n%3 == 0:
		s = "Fizz"
	case n%5 == 0:
		s = "Buzz"
	default:
		s = fmt.Sprintf("%d", n)
	}
	return
}
