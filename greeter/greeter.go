// Using Test first principles, implement a set of tests for your very own
// greeter function that we've written in the fmt lab.
//
// o Use ad-hoc values for name, age, height
// o Try out zero-values for name, age, height
// o Make use of the assert functions from testify to validate your tests.
// o See installation/run steps below
// o Practice breaking one of the test on purpose and fixing it.
// o Issue a test command to only run one of the tests from the command line!
// o Make sure you achieve 100% test coverage using the standard coverage tool.

package greeter

import "fmt"

// greets a fellow human
const greet = "Hello, my name is %s. I am %b years old and '%05.2f' tall!"

// greeter return a greeting based on name, age and height
func greeter(n string, a int, h float32) string {
	return fmt.Sprintf(greet, n, a, h)
}
