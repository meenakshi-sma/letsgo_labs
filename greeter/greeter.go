// Greeter computes a greeting message

package greeter

import (
	"fmt"
	"strconv"
)

const greeting = "Hello, %s! You are %d old today..."

// greeter return a greeting based on name, age and height
func greeter(n string, a int) string {
	return fmt.Sprintf(greeting, n, a)
}

// greeter1 return a greeting based on name, age and height
func greeter1(n string, a int) string {
	return "Hello, " + n + "! You are " + strconv.Itoa(a) + " old today..."
}
