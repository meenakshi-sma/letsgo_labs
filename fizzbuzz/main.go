package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 0, "Enter a number")
	flag.Parse()

	for i := 1; i <= *n; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
