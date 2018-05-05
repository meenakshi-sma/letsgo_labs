package greeter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeter(t *testing.T) {
	n, a := "Fernand", 30

	assert.Equal(t, fmt.Sprintf(greeting, n, a), greeter(n, a))
}

func TestGreeter1(t *testing.T) {
	n, a := "Fernand", 30

	assert.Equal(t, fmt.Sprintf(greeting, n, a), greeter1(n, a))
}

func TestGreeter2(t *testing.T) {
	n, a := "Fernand", 30

	assert.Equal(t, fmt.Sprintf(greeting, n, a), greeter2(n, a))
}

func TestGreeter3(t *testing.T) {
	n, a := "Fernand", 30

	assert.Equal(t, fmt.Sprintf(greeting, n, a), greeter3(n, a))
}

func BenchmarkGreeter(b *testing.B) {
	n, a := "Fernand", 30

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = greeter(n, a)
	}
}

func BenchmarkGreeter1(b *testing.B) {
	n, a := "Fernand", 30

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = greeter1(n, a)
	}
}

func BenchmarkGreeter2(b *testing.B) {
	n, a := "Fernand", 30

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = greeter2(n, a)
	}
}

func BenchmarkGreeter3(b *testing.B) {
	n, a := "Fernand", 30

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = greeter3(n, a)
	}
}
