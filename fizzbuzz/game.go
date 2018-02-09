// © 2018 Imhotep Software LLC. All rights reserved.

// Package fizzbuzz implements a FizzBuzz game.
//
// In a FizzBuzz Game, children count up from 1.
// 	o If the number is a multiple of 3, they say **Fizz**
// 	o If the number is a multiple of 5, they say **Buzz**
// 	o For multiple of both 3 and 5 they say **FizzBuzz**
// 	o Otherwise they say the number
package fizzbuzz

import (
	"fmt"
)

// Compute a FizzBuzz number based on given input.
//
// 	Returns `FizzBuzz if number div by 3 and 5.
// 	`Fizz if div by 3.
// 	`Buzz if div by 5.
// 	`number otherwise.
func Compute(n int) (s string) {
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
