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

// greeter return a greeting based on name, age and height
func greeter1(n string, a int, h float32) string {
	gret=fmt.snprintf( "Hi %s welcome", n, a, h)
	
}

// greeter return a greeting based on name, age and height
func greeter2(n string, a int, h float32) string {
	gret=fmt.snprintf( "Hi %s welcome on your height",n, a, h)
}
