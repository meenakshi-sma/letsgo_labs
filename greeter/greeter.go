// Greeter computes a greeting message

package greeter

import (
	"fmt"
	"strconv"
	"strings"
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

func greeter2(n string, a int) string {
	var b strings.Builder

	b.Grow(39)
	// fmt.Fprintf(&b, "Hello, "+n+"! You are "+strconv.Itoa(a)+" old today...")
	b.WriteString("Hello, ")
	b.WriteString(n)
	b.WriteString("! You are ")
	b.WriteString(strconv.Itoa(a))
	b.WriteString(" old today...")

	return b.String()
}

func greeter3(n string, a int) string {
	var buff = make([]byte, 39)

	copy(buff, "Hello, "+n+"! You are "+strconv.Itoa(a)+" old today...")

	return string(buff)
}
