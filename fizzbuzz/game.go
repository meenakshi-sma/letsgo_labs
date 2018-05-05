// © 2018 Imhotep Software LLC. All rights reserved.

// Package fizzbuzz implements a FizzBuzz game.
//
// In a FizzBuzz Game, children count up from 1.
// 	o If the number is a multiple of 3, they say **Fizz**
// 	o If the number is a multiple of 5, they say **Buzz**
// 	o For multiple of both 3 and 5 they say **FizzBuzz**
// 	o Otherwise they say the number
package fizzbuzz

import "strconv"

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzbuzz = "FizzBuzz"
)

// Compute a FizzBuzz number based on given input.
//
// 	Returns `FizzBuzz if number div by 3 and 5.
// 	`Fizz if div by 3.
// 	`Buzz if div by 5.
// 	`number otherwise.
func compute(n int) (s string) {
	switch {
	case n%3 == 0 && n%5 == 0:
		s = fizzbuzz
	case n%3 == 0:
		s = fizz
	case n%5 == 0:
		s = buzz
	default:
		s = strconv.Itoa(n)
	}
	return
}
