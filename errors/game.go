package main

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzBuzz = fizz + buzz
)

var (
	errUnderRange = errors.New("number is under range (<=0)")
	errOverRange  = errors.New("number is over range (> 20)")
)

func main() {
	for i := 0; i <= 21; i++ {
		fmt.Printf("%02d ", i)
		if r, err := play(i); err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", r)
		}
	}
}

func play(n int) (string, error) {
	var s string

	if n <= 0 {
		return s, errors.Wrapf(errUnderRange, "FizzBuzz with %d", n)
	}

	if n > 20 {
		return s, errors.Wrapf(errOverRange, "FizzBuzz with %d", n)
	}

	switch {
	case n%3 == 0 && n%5 == 0:
		s = fizzBuzz
	case n%3 == 0:
		s = fizz
	case n%5 == 0:
		s = buzz
	default:
		s = fmt.Sprintf("%d", n)
	}

	return s, nil
}
