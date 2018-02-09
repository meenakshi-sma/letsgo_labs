package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		name string
		age  int
	)

	flag.StringVar(&name, "u", "No one", "specifies a user name. Default: No one")
	flag.IntVar(&age, "a", 42, "specifies user age. Default: 42")

	// Don't Forget!
	flag.Parse()

	fmt.Printf("Hello my name is %s and I am %d years old!\n", name, age)
}
