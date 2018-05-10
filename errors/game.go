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

// Define errors...
// YOUR_CODE...


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

func play(n int) // YOUR_CODE {
  // YOUR_CODE...
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
	// YOUR_CODE...
}
