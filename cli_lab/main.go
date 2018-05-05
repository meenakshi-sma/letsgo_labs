package main

import (
	"flag"
	"fmt"
)

func greet(name string, age int) string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old!", name, age)
}

func main() {
	u := flag.String("u", "No One", "Specify a user name")
	a := flag.Int("a", 42, "Specify a user age")
	flag.Parse()

	fmt.Println(greet(*u, *a))
}
