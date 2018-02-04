package main

import (
	"fmt"
)

/*
FizzBuzz returns `FizzBuzz` is number div by 3 and 5.
 Fizz is div by 3.
 Buzz if div by 5 or number otherwise
*/
func FizzBuzz(n int) (s string) {
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
